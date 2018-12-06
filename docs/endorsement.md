# 背书策略

## (常用)策略类型简介

> 如下图
![策略结构](images/policy.png)
如上图中,目前只实现了 SignaturePolicy 和 ImplicitMetaPolicy 两种策略;  
两种策略数据结构参考《背书策略.vsdx》文档中的两种相应策略的页签;

## sdk背书策略接口实现

### 任意一个成员签名

> 成员:(目前理解)包括msp中的peer节点实体，client用户等;
> 选取组织[ids]中任意一个MSPID中的任意一个Member签名,即可满足策略.
> msp角色是:MSPRole_MEMBER

    ```
    function:
        SignedByAnyMember(ids []string) *cb.SignaturePolicyEnvelope {...}
    params：
        ids [组织MSPID]]
    return:
        SignaturePolicyEnvelope 签名背书策略Envelop

    eg:
       选取三个组织(Org1MSP、Org2MSP、Org3MSP)中任意一个成员Member背书即可满足策略.
       用OR表达式表示为：
          背书策略为:OR ["Org1MSP.Member","Org2MSP.Member","Org3MSP.Member"]
       用此函数表示为:
          var ids []string{"Org1MSP","Org2MSP","Org3MSP"}
          endorsementplicy := SignedByAnyMember(ids)
          endorsementplicy就是执行chaincode实例化时的背书策略.
    ```

### 任意一个peer实体签名

> 选取组织[ids]中的任意一个 peer节点签名,即可满足策略.
> msp角色是:MSPRole_PEER
> 例如：Org1MSP.Peer

    ```
    function:
        SignedByAnyPeer(ids []string) *cb.SignaturePolicyEnvelope {...}
    params：
        ids [组织MSPID]]
    return:
        SignaturePolicyEnvelope 签名背书策略Envelop

    eg:
    选取三个组织(Org1MSP、Org2MSP、Org3MSP)中任意一个peer节点背书即可满足策略.
    用OR表达式表示为：
        背书策略为:OR ["Org1MSP.Peer","Org2MSP.Peer","Org3MSP.Peer"]
    用此函数表示为:
        var ids []string{"Org1MSP","Org2MSP","Org3MSP"}
        endorsementplicy := SignedByAnyMember(ids)
        endorsementplicy就是执行chaincode实例化时的背书策略.
    ```

### MP1 || (MP2 && MP3)

    ```
    function:
        SignedByGivenRoleMP(ids []string) *cb.SignaturePolicyEnvelope {...}
        具体实现参考下面.
    params：
        mspids [组织MSPID]] [Org1MSP,Org2MSP,Org3MSP]
    return:
        SignaturePolicyEnvelope 签名背书策略Envelop

    ```

    ```
    eg：实现 MP1 || (MP2 && MP3)
        // MP1 || (MP2 && MP3)
        func SignedByGivenRoleMP(ids []string) *cb.SignaturePolicyEnvelope {
            return signedByGivenRole(msp.MSPRole_MEMBER, ids)
        }

        // role = "MSPRole_MEMBER"
        func signedByGivenRole(role msp.MSPRole_MSPRoleType, ids []string) *cb.SignaturePolicyEnvelope {

            sort.Strings(ids)
            principals := make([]*msp.MSPPrincipal, len(ids))

            signPolicyAnd := make([]*cb.SignaturePolicy,2)
            signPolicyAnd[0] = SignedBy(int32(1))
            signPolicyAnd[1] = SignedBy(int32(2))

            sigspolicy := make([]*cb.SignaturePolicy, 2)
            sigspolicy[0] = SignedBy(int32(0))
            sigspolicy[1] = NOutOf(2, signPolicyAnd)

            for i, id := range ids {
                principals[i] = &msp.MSPPrincipal{
                    PrincipalClassification: msp.MSPPrincipal_ROLE,
                    Principal:               utils.MarshalOrPanic(&msp.MSPRole{Role: role, MspIdentifier: id}),
                }
            }

            p := &cb.SignaturePolicyEnvelope{
                Version:    0,
                Rule:       NOutOf(1, sigspolicy),
                Identities: principals,
            }

            return p
        }

    ```



### 通过指定的组织的Member签名

    ```
    function:
        SignedByAssignMember(n int32 , mspids []string) *cb.SignaturePolicyEnvelope {...}
        具体实现参考下面.
    params：
        n ：n < len(mspids) 从mspids中选取n个组织(从这n个组织中选取节点进行背书)
        mspids [组织MSPID]]
    return:
        SignaturePolicyEnvelope 签名背书策略Envelop

    eg:
     选取三个组织(Org1MSP、Org2MSP、Org3MSP)中任意2个组织的peer节点进行背书:
    用此函数表示为:
        endorsementplicy := SignedByAssignMember(2,[]string{"Org1MSP","Org2MSP","Org3MSP"})
        endorsementplicy 就是执行chaincode实例化时的背书策略.
    ```

    ```
    func SignedByAssignMember(n int32 , ids []string) *cb.SignaturePolicyEnvelope {
        return signedByAssignOfGivenRole(n , msp.MSPRole_MEMBER, ids)
    }

    func signedByAssignOfGivenRole(n int32, role msp.MSPRole_MSPRoleType, ids []string) *cb.SignaturePolicyEnvelope {
        sort.Strings(ids)
        principals := make([]*msp.MSPPrincipal, len(ids))
        sigspolicy := make([]*cb.SignaturePolicy, len(ids))
        for i, id := range ids {
            principals[i] = &msp.MSPPrincipal{
                PrincipalClassification: msp.MSPPrincipal_ROLE,
                Principal:               utils.MarshalOrPanic(&msp.MSPRole{Role: role, MspIdentifier: id})}
            sigspolicy[i] = cauthdsl.SignedBy(int32(i))
        }

        // create the policy: it requires exactly 1 signature from any of the principals
        p := &cb.SignaturePolicyEnvelope{
            Version:    0,
            Rule:       cauthdsl.NOutOf(n, sigspolicy),
            Identities: principals,
        }

        return p
    }

    ```

