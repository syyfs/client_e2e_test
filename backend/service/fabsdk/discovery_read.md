# 服务发现

> 具体实现示例参考client_e2e_test中的DiscoverService
> 服务发现提供了3个接口:  

```bash
   1. peer节点信息查询
        获取的是某channel的区块链网络中的所有peer节点信息;
        可以用两种方式获取:(分别在下面有具体示例)
        利用服务发现服务,调用以下函数:
        discoveryService.GetPeers()  
        or  
        AddPeersQuery()
    2. 获取某channel配置信息
        利用服务发现服务,调用以下函数:
        AddConfigQuery()
    3. 背书节点信息查询
        利用服务发现服务,调用以下函数:
        AddEndorsersQuery()
```

## step1.准备服务发现Factory步骤

```bash
  1. 因需引用factory的ProviderFactory和ChannelProvider，因此需要重新定义结构体来引用ProviderFactory和ChannelProvider:
    type DynamicDiscoveryProviderFactory struct {  
        // 外部引用
        defsvc.ProviderFactory
    }
    type channelProvider struct {
        fab.ChannelProvider
        services map[string]*dynamicdiscovery.ChannelService
    }
    type channelService struct {
        fab.ChannelService
        discovery fab.DiscoveryService
    }
  2. 重写DynamicDiscoveryProviderFactory中的CreateChannelProvider函数;
  3. 重写channelProvider中的 ChannelService() 以及 Discovery()函数;
```

## peer节点信息查询

- peer节点信息查询，输出结果  

```bash

[
    {
        "MSPID": "Org2MSP",
        "LedgerHeight": 5,
        "Endpoint": "peer0.org2.example.com:7051",
        "Identity": "-----BEGIN CERTIFICATE-----\nMIICKTCCAc+gAwIBAgIRANK4WBck5gKuzTxVQIwhYMUwCgYIKoZIzj0EAwIwczEL\nMAkGA1UEBhMCVVMxEzARBgNVBAgTCkNhbGlmb3JuaWExFjAUBgNVBAcTDVNhbiBG\ncmFuY2lzY28xGTAXBgNVBAoTEG9yZzIuZXhhbXBsZS5jb20xHDAaBgNVBAMTE2Nh\nLm9yZzIuZXhhbXBsZS5jb20wHhcNMTgwNjE3MTM0NTIxWhcNMjgwNjE0MTM0NTIx\nWjBqMQswCQYDVQQGEwJVUzETMBEGA1UECBMKQ2FsaWZvcm5pYTEWMBQGA1UEBxMN\nU2FuIEZyYW5jaXNjbzENMAsGA1UECxMEcGVlcjEfMB0GA1UEAxMWcGVlcjAub3Jn\nMi5leGFtcGxlLmNvbTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABJa0gkMRqJCi\nzmx+L9xy/ecJNvdAV2zmSx5Sf2qospVAH1MYCHyudDEvkiRuBPgmCdOdwJsE0g+h\nz0nZdKq6/X+jTTBLMA4GA1UdDwEB/wQEAwIHgDAMBgNVHRMBAf8EAjAAMCsGA1Ud\nIwQkMCKAIFZMuZfUtY6n2iyxaVr3rl+x5lU0CdG9x7KAeYydQGTMMAoGCCqGSM49\nBAMCA0gAMEUCIQC0M9/LJ7j3I9NEPQ/B1BpnJP+UNPnGO2peVrM/mJ1nVgIgS1ZA\nA1tsxuDyllaQuHx2P+P9NDFdjXx5T08lZhxuWYM=\n-----END CERTIFICATE-----\n",
        "Chaincodes": [
            "mycc"
        ]
    },
]
```

- 代码示例一、  

```bash
    1. 获取FabricSDK实例并获取ChannelProvider
       sdk , err := fabsdk.New(configProvider,fabsdk.WithServicePkg(&DynamicDiscoveryProviderFactory{}))
       require.NoError(t, err, "Failed to create new SDK")
       chProvider := sdk.ChannelContext("mychannel", fabsdk.WithUser("Admin"), fabsdk.WithOrg("Org1"))
       chCtx, err := chProvider()
    2. 使用 ChannelProvider 获取discoverservice
       discoveryService, err := chCtx.ChannelService().Discovery()

    3. 使用接口获取网络中所有的peer节点
        peers, err := discoveryService.GetPeers()
        for _, p := range peers {
            fmt.Printf("- [%s] - MSP [%s] \n ", p.URL(), p.MSPID())
        }
```

- 代码示例二、  

> discovery client Request AddPeersQuery() --> Responses Get Peers()

```bash
    1.  获取FabricSDK实例并获取ClientProvider
        sdk , err := fabsdk.New(configProvider,fabsdk.WithServicePkg(&DynamicDiscoveryProviderFactory{}))
        require.NoError(t, err, "Failed to create new SDK")
        defer sdk.Close()
        chProvider := sdk.Context(fabsdk.WithUser("Admin"), fabsdk.WithOrg("Org1"))
        chCtx, err := chProvider()
    2. 创建服务发现客户端discovery.Client请求，并添加PeersQuery；向某targetpeer 发送查询请求
        req := discclient.NewRequest().OfChannel(orgChannelID).AddPeersQuery()
        responses, err := client.Send(reqCtx, req,peerCfg1.PeerConfig)
    3. 获取某channel中的所有peers，可以循环输出peers节点信息
        resp := responses[0]
        chanResp := resp.ForChannel(orgChannelID)
        peers, err := chanResp.Peers()
```

## 获取某channel配置信息

- 获取某channel配置信息，输出结果  

```bash
{
    "msps": {
        "OrdererOrg": {
            "name": "OrdererMSP",
            "root_certs": [
                "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUNMekNDQWRhZ0F3SUJBZ0lSQU1pWkxUb3RmMHR6VTRzNUdIdkQ0UjR3Q2dZSUtvWkl6ajBFQXdJd2FURUwKTUFrR0ExVUVCaE1DVlZNeEV6QVJCZ05WQkFnVENrTmhiR2xtYjNKdWFXRXhGakFVQmdOVkJBY1REVk5oYmlCRwpjbUZ1WTJselkyOHhGREFTQmdOVkJBb1RDMlY0WVcxd2JHVXVZMjl0TVJjd0ZRWURWUVFERXc1allTNWxlR0Z0CmNHeGxMbU52YlRBZUZ3MHhPREEyTURreE1UVTRNamhhRncweU9EQTJNRFl4TVRVNE1qaGFNR2t4Q3pBSkJnTlYKQkFZVEFsVlRNUk13RVFZRFZRUUlFd3BEWVd4cFptOXlibWxoTVJZd0ZBWURWUVFIRXcxVFlXNGdSbkpoYm1OcApjMk52TVJRd0VnWURWUVFLRXd0bGVHRnRjR3hsTG1OdmJURVhNQlVHQTFVRUF4TU9ZMkV1WlhoaGJYQnNaUzVqCmIyMHdXVEFUQmdjcWhrak9QUUlCQmdncWhrak9QUU1CQndOQ0FBUW9ySjVSamFTQUZRci9yc2xoMWdobnNCWEQKeDVsR1lXTUtFS1pDYXJDdkZBekE0bHUwb2NQd0IzNWJmTVN5bFJPVmdVdHF1ZU9IcFBNc2ZLNEFrWjR5bzE4dwpYVEFPQmdOVkhROEJBZjhFQkFNQ0FhWXdEd1lEVlIwbEJBZ3dCZ1lFVlIwbEFEQVBCZ05WSFJNQkFmOEVCVEFECkFRSC9NQ2tHQTFVZERnUWlCQ0JnbmZJd0pzNlBaWUZCclpZVkRpU05vSjNGZWNFWHYvN2xHL3QxVUJDbVREQUsKQmdncWhrak9QUVFEQWdOSEFEQkVBaUE5NGFkc21UK0hLalpFVVpnM0VkaWdSM296L3pEQkNhWUY3TEJUVXpuQgpEZ0lnYS9RZFNPQnk1TUx2c0lSNTFDN0N4UnR2NUM5V05WRVlmWk5SaGdXRXpoOD0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="
            ],
            "admins": [
                "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUNDVENDQWJDZ0F3SUJBZ0lRR2wzTjhaSzRDekRRQmZqYVpwMVF5VEFLQmdncWhrak9QUVFEQWpCcE1Rc3cKQ1FZRFZRUUdFd0pWVXpFVE1CRUdBMVVFQ0JNS1EyRnNhV1p2Y201cFlURVdNQlFHQTFVRUJ4TU5VMkZ1SUVaeQpZVzVqYVhOamJ6RVVNQklHQTFVRUNoTUxaWGhoYlhCc1pTNWpiMjB4RnpBVkJnTlZCQU1URG1OaExtVjRZVzF3CmJHVXVZMjl0TUI0WERURTRNRFl3T1RFeE5UZ3lPRm9YRFRJNE1EWXdOakV4TlRneU9Gb3dWakVMTUFrR0ExVUUKQmhNQ1ZWTXhFekFSQmdOVkJBZ1RDa05oYkdsbWIzSnVhV0V4RmpBVUJnTlZCQWNURFZOaGJpQkdjbUZ1WTJsegpZMjh4R2pBWUJnTlZCQU1NRVVGa2JXbHVRR1Y0WVcxd2JHVXVZMjl0TUZrd0V3WUhLb1pJemowQ0FRWUlLb1pJCnpqMERBUWNEUWdBRWl2TXQybVdiQ2FHb1FZaWpka1BRM1NuTGFkMi8rV0FESEFYMnRGNWthMTBteG1OMEx3VysKdmE5U1dLMmJhRGY5RDQ2TVROZ2gycnRhUitNWXFWRm84Nk5OTUVzd0RnWURWUjBQQVFIL0JBUURBZ2VBTUF3RwpBMVVkRXdFQi93UUNNQUF3S3dZRFZSMGpCQ1F3SW9BZ1lKM3lNQ2JPajJXQlFhMldGUTRramFDZHhYbkJGNy8rCjVSdjdkVkFRcGt3d0NnWUlLb1pJemowRUF3SURSd0F3UkFJZ2RIc0pUcGM5T01DZ3JPVFRLTFNnU043UWk3MWIKSWpkdzE4MzJOeXFQZnJ3Q0lCOXBhSlRnL2R5ckNhWUx1ZndUbUtFSnZZMEtXVzcrRnJTeG5CTGdzZjJpCi0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K"
            ],
            "crypto_config": {
                "signature_hash_family": "SHA2",
                "identity_identifier_hash_function": "SHA256"
            },
            "tls_root_certs": [
                "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUNORENDQWR1Z0F3SUJBZ0lRZDdodzFIaHNZTXI2a25ETWJrZThTakFLQmdncWhrak9QUVFEQWpCc01Rc3cKQ1FZRFZRUUdFd0pWVXpFVE1CRUdBMVVFQ0JNS1EyRnNhV1p2Y201cFlURVdNQlFHQTFVRUJ4TU5VMkZ1SUVaeQpZVzVqYVhOamJ6RVVNQklHQTFVRUNoTUxaWGhoYlhCc1pTNWpiMjB4R2pBWUJnTlZCQU1URVhSc2MyTmhMbVY0CllXMXdiR1V1WTI5dE1CNFhEVEU0TURZd09URXhOVGd5T0ZvWERUSTRNRFl3TmpFeE5UZ3lPRm93YkRFTE1Ba0cKQTFVRUJoTUNWVk14RXpBUkJnTlZCQWdUQ2tOaGJHbG1iM0p1YVdFeEZqQVVCZ05WQkFjVERWTmhiaUJHY21GdQpZMmx6WTI4eEZEQVNCZ05WQkFvVEMyVjRZVzF3YkdVdVkyOXRNUm93R0FZRFZRUURFeEYwYkhOallTNWxlR0Z0CmNHeGxMbU52YlRCWk1CTUdCeXFHU000OUFnRUdDQ3FHU000OUF3RUhBMElBQk9ZZGdpNm53a3pYcTBKQUF2cTIKZU5xNE5Ybi85L0VRaU13Tzc1dXdpTWJVbklYOGM1N2NYU2dQdy9NMUNVUGFwNmRyMldvTjA3RGhHb1B6ZXZaMwp1aFdqWHpCZE1BNEdBMVVkRHdFQi93UUVBd0lCcGpBUEJnTlZIU1VFQ0RBR0JnUlZIU1VBTUE4R0ExVWRFd0VCCi93UUZNQU1CQWY4d0tRWURWUjBPQkNJRUlCcW0xZW9aZy9qSW52Z1ZYR2cwbzVNamxrd2tSekRlalAzZkplbW8KU1hBek1Bb0dDQ3FHU000OUJBTUNBMGNBTUVRQ0lEUG9FRkF5bFVYcEJOMnh4VEo0MVplaS9ZQWFvN29aL0tEMwpvTVBpQ3RTOUFpQmFiU1dNS3UwR1l4eXdsZkFwdi9CWitxUEJNS0JMNk5EQ1haUnpZZmtENEE9PQotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg=="
            ]
        },
       "Org1MSP": {
            "name": "Org1MSP",
            "root_certs": [
                "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUNSRENDQWVxZ0F3SUJBZ0lSQU1nN2VETnhwS0t0ZGl0TDRVNDRZMUl3Q2dZSUtvWkl6ajBFQXdJd2N6RUwKTUFrR0ExVUVCaE1DVlZNeEV6QVJCZ05WQkFnVENrTmhiR2xtYjNKdWFXRXhGakFVQmdOVkJBY1REVk5oYmlCRwpjbUZ1WTJselkyOHhHVEFYQmdOVkJBb1RFRzl5WnpFdVpYaGhiWEJzWlM1amIyMHhIREFhQmdOVkJBTVRFMk5oCkxtOXlaekV1WlhoaGJYQnNaUzVqYjIwd0hoY05NVGd3TmpBNU1URTFPREk0V2hjTk1qZ3dOakEyTVRFMU9ESTQKV2pCek1Rc3dDUVlEVlFRR0V3SlZVekVUTUJFR0ExVUVDQk1LUTJGc2FXWnZjbTVwWVRFV01CUUdBMVVFQnhNTgpVMkZ1SUVaeVlXNWphWE5qYnpFWk1CY0dBMVVFQ2hNUWIzSm5NUzVsZUdGdGNHeGxMbU52YlRFY01Cb0dBMVVFCkF4TVRZMkV1YjNKbk1TNWxlR0Z0Y0d4bExtTnZiVEJaTUJNR0J5cUdTTTQ5QWdFR0NDcUdTTTQ5QXdFSEEwSUEKQk41d040THpVNGRpcUZSWnB6d3FSVm9JbWw1MVh0YWkzbWgzUXo0UEZxWkhXTW9lZ0ovUWRNKzF4L3RobERPcwpnbmVRcndGd216WGpvSSszaHJUSmRuU2pYekJkTUE0R0ExVWREd0VCL3dRRUF3SUJwakFQQmdOVkhTVUVDREFHCkJnUlZIU1VBTUE4R0ExVWRFd0VCL3dRRk1BTUJBZjh3S1FZRFZSME9CQ0lFSU9CZFFMRitjTVdhNmUxcDJDcE8KRXg3U0hVaW56VnZkNTVoTG03dzZ2NzJvTUFvR0NDcUdTTTQ5QkFNQ0EwZ0FNRVVDSVFDQyt6T1lHcll0ZTB4SgpSbDVYdUxjUWJySW9UeHpsRnJLZWFNWnJXMnVaSkFJZ0NVVGU5MEl4aW55dk4wUkh4UFhoVGNJTFdEZzdLUEJOCmVrNW5TRlh3Y0lZPQotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg=="
            ],
            "admins": [
                "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUNLakNDQWRDZ0F3SUJBZ0lRRTRFK0tqSHgwdTlzRSsxZUgrL1dOakFLQmdncWhrak9QUVFEQWpCek1Rc3cKQ1FZRFZRUUdFd0pWVXpFVE1CRUdBMVVFQ0JNS1EyRnNhV1p2Y201cFlURVdNQlFHQTFVRUJ4TU5VMkZ1SUVaeQpZVzVqYVhOamJ6RVpNQmNHQTFVRUNoTVFiM0puTVM1bGVHRnRjR3hsTG1OdmJURWNNQm9HQTFVRUF4TVRZMkV1CmIzSm5NUzVsZUdGdGNHeGxMbU52YlRBZUZ3MHhPREEyTURreE1UVTRNamhhRncweU9EQTJNRFl4TVRVNE1qaGEKTUd3eEN6QUpCZ05WQkFZVEFsVlRNUk13RVFZRFZRUUlFd3BEWVd4cFptOXlibWxoTVJZd0ZBWURWUVFIRXcxVApZVzRnUm5KaGJtTnBjMk52TVE4d0RRWURWUVFMRXdaamJHbGxiblF4SHpBZEJnTlZCQU1NRmtGa2JXbHVRRzl5Clp6RXVaWGhoYlhCc1pTNWpiMjB3V1RBVEJnY3Foa2pPUFFJQkJnZ3Foa2pPUFFNQkJ3TkNBQVFqK01MZk1ESnUKQ2FlWjV5TDR2TnczaWp4ZUxjd2YwSHo1blFrbXVpSnFETjRhQ0ZwVitNTTVablFEQmx1dWRyUS80UFA1Sk1WeQpreWZsQ3pJa2NCNjdvMDB3U3pBT0JnTlZIUThCQWY4RUJBTUNCNEF3REFZRFZSMFRBUUgvQkFJd0FEQXJCZ05WCkhTTUVKREFpZ0NEZ1hVQ3hmbkRGbXVudGFkZ3FUaE1lMGgxSXA4MWIzZWVZUzV1OE9yKzlxREFLQmdncWhrak8KUFFRREFnTklBREJGQWlFQXlJV21QcjlQakdpSk1QM1pVd05MRENnNnVwMlVQVXNJSzd2L2h3RVRra01DSUE0cQo3cHhQZy9VVldiamZYeE0wUCsvcTEzbXFFaFlYaVpTTXpoUENFNkNmCi0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K"
            ],
            "crypto_config": {
                "signature_hash_family": "SHA2",
                "identity_identifier_hash_function": "SHA256"
            },
            "tls_root_certs": [
                "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUNTVENDQWUrZ0F3SUJBZ0lRZlRWTE9iTENVUjdxVEY3Z283UXgvakFLQmdncWhrak9QUVFEQWpCMk1Rc3cKQ1FZRFZRUUdFd0pWVXpFVE1CRUdBMVVFQ0JNS1EyRnNhV1p2Y201cFlURVdNQlFHQTFVRUJ4TU5VMkZ1SUVaeQpZVzVqYVhOamJ6RVpNQmNHQTFVRUNoTVFiM0puTVM1bGVHRnRjR3hsTG1OdmJURWZNQjBHQTFVRUF4TVdkR3h6ClkyRXViM0puTVM1bGVHRnRjR3hsTG1OdmJUQWVGdzB4T0RBMk1Ea3hNVFU0TWpoYUZ3MHlPREEyTURZeE1UVTQKTWpoYU1IWXhDekFKQmdOVkJBWVRBbFZUTVJNd0VRWURWUVFJRXdwRFlXeHBabTl5Ym1saE1SWXdGQVlEVlFRSApFdzFUWVc0Z1JuSmhibU5wYzJOdk1Sa3dGd1lEVlFRS0V4QnZjbWN4TG1WNFlXMXdiR1V1WTI5dE1SOHdIUVlEClZRUURFeFowYkhOallTNXZjbWN4TG1WNFlXMXdiR1V1WTI5dE1Ga3dFd1lIS29aSXpqMENBUVlJS29aSXpqMEQKQVFjRFFnQUVZbnp4bmMzVUpHS0ZLWDNUNmR0VGpkZnhJTVYybGhTVzNab0lWSW9mb04rWnNsWWp0d0g2ZXZXYgptTkZGQmRaYWExTjluaXRpbmxxbVVzTU1NQ2JieXFOZk1GMHdEZ1lEVlIwUEFRSC9CQVFEQWdHbU1BOEdBMVVkCkpRUUlNQVlHQkZVZEpRQXdEd1lEVlIwVEFRSC9CQVV3QXdFQi96QXBCZ05WSFE0RUlnUWdlVTAwNlNaUllUNDIKN1Uxb2YwL3RGdHUvRFVtazVLY3hnajFCaklJakduZ3dDZ1lJS29aSXpqMEVBd0lEU0FBd1JRSWhBSWpvcldJTwpRNVNjYjNoZDluRi9UamxWcmk1UHdTaDNVNmJaMFdYWEsxYzVBaUFlMmM5QmkyNFE1WjQ0aXQ1MkI5cm1hU1NpCkttM2NZVlY0cWJ6RFhMOHZYUT09Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K"
            ],
            "fabric_node_ous": {
                "enable": true,
                "client_ou_identifier": {
                    "certificate": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUNSRENDQWVxZ0F3SUJBZ0lSQU1nN2VETnhwS0t0ZGl0TDRVNDRZMUl3Q2dZSUtvWkl6ajBFQXdJd2N6RUwKTUFrR0ExVUVCaE1DVlZNeEV6QVJCZ05WQkFnVENrTmhiR2xtYjNKdWFXRXhGakFVQmdOVkJBY1REVk5oYmlCRwpjbUZ1WTJselkyOHhHVEFYQmdOVkJBb1RFRzl5WnpFdVpYaGhiWEJzWlM1amIyMHhIREFhQmdOVkJBTVRFMk5oCkxtOXlaekV1WlhoaGJYQnNaUzVqYjIwd0hoY05NVGd3TmpBNU1URTFPREk0V2hjTk1qZ3dOakEyTVRFMU9ESTQKV2pCek1Rc3dDUVlEVlFRR0V3SlZVekVUTUJFR0ExVUVDQk1LUTJGc2FXWnZjbTVwWVRFV01CUUdBMVVFQnhNTgpVMkZ1SUVaeVlXNWphWE5qYnpFWk1CY0dBMVVFQ2hNUWIzSm5NUzVsZUdGdGNHeGxMbU52YlRFY01Cb0dBMVVFCkF4TVRZMkV1YjNKbk1TNWxlR0Z0Y0d4bExtTnZiVEJaTUJNR0J5cUdTTTQ5QWdFR0NDcUdTTTQ5QXdFSEEwSUEKQk41d040THpVNGRpcUZSWnB6d3FSVm9JbWw1MVh0YWkzbWgzUXo0UEZxWkhXTW9lZ0ovUWRNKzF4L3RobERPcwpnbmVRcndGd216WGpvSSszaHJUSmRuU2pYekJkTUE0R0ExVWREd0VCL3dRRUF3SUJwakFQQmdOVkhTVUVDREFHCkJnUlZIU1VBTUE4R0ExVWRFd0VCL3dRRk1BTUJBZjh3S1FZRFZSME9CQ0lFSU9CZFFMRitjTVdhNmUxcDJDcE8KRXg3U0hVaW56VnZkNTVoTG03dzZ2NzJvTUFvR0NDcUdTTTQ5QkFNQ0EwZ0FNRVVDSVFDQyt6T1lHcll0ZTB4SgpSbDVYdUxjUWJySW9UeHpsRnJLZWFNWnJXMnVaSkFJZ0NVVGU5MEl4aW55dk4wUkh4UFhoVGNJTFdEZzdLUEJOCmVrNW5TRlh3Y0lZPQotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg==",
                    "organizational_unit_identifier": "client"
                },
                "peer_ou_identifier": {
                    "certificate": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUNSRENDQWVxZ0F3SUJBZ0lSQU1nN2VETnhwS0t0ZGl0TDRVNDRZMUl3Q2dZSUtvWkl6ajBFQXdJd2N6RUwKTUFrR0ExVUVCaE1DVlZNeEV6QVJCZ05WQkFnVENrTmhiR2xtYjNKdWFXRXhGakFVQmdOVkJBY1REVk5oYmlCRwpjbUZ1WTJselkyOHhHVEFYQmdOVkJBb1RFRzl5WnpFdVpYaGhiWEJzWlM1amIyMHhIREFhQmdOVkJBTVRFMk5oCkxtOXlaekV1WlhoaGJYQnNaUzVqYjIwd0hoY05NVGd3TmpBNU1URTFPREk0V2hjTk1qZ3dOakEyTVRFMU9ESTQKV2pCek1Rc3dDUVlEVlFRR0V3SlZVekVUTUJFR0ExVUVDQk1LUTJGc2FXWnZjbTVwWVRFV01CUUdBMVVFQnhNTgpVMkZ1SUVaeVlXNWphWE5qYnpFWk1CY0dBMVVFQ2hNUWIzSm5NUzVsZUdGdGNHeGxMbU52YlRFY01Cb0dBMVVFCkF4TVRZMkV1YjNKbk1TNWxlR0Z0Y0d4bExtTnZiVEJaTUJNR0J5cUdTTTQ5QWdFR0NDcUdTTTQ5QXdFSEEwSUEKQk41d040THpVNGRpcUZSWnB6d3FSVm9JbWw1MVh0YWkzbWgzUXo0UEZxWkhXTW9lZ0ovUWRNKzF4L3RobERPcwpnbmVRcndGd216WGpvSSszaHJUSmRuU2pYekJkTUE0R0ExVWREd0VCL3dRRUF3SUJwakFQQmdOVkhTVUVDREFHCkJnUlZIU1VBTUE4R0ExVWRFd0VCL3dRRk1BTUJBZjh3S1FZRFZSME9CQ0lFSU9CZFFMRitjTVdhNmUxcDJDcE8KRXg3U0hVaW56VnZkNTVoTG03dzZ2NzJvTUFvR0NDcUdTTTQ5QkFNQ0EwZ0FNRVVDSVFDQyt6T1lHcll0ZTB4SgpSbDVYdUxjUWJySW9UeHpsRnJLZWFNWnJXMnVaSkFJZ0NVVGU5MEl4aW55dk4wUkh4UFhoVGNJTFdEZzdLUEJOCmVrNW5TRlh3Y0lZPQotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg==",
                    "organizational_unit_identifier": "peer"
                }
            }
        },
    }
}
```

- 代码示例  
> discovery client Request AddConfigQuery() --> Responses Get Config()  

```bash
    1.  获取FabricSDK实例并获取ClientProvider
        sdk , err := fabsdk.New(configProvider,fabsdk.WithServicePkg(&DynamicDiscoveryProviderFactory{}))
        require.NoError(t, err, "Failed to create new SDK")
        defer sdk.Close()
        chProvider := sdk.Context(fabsdk.WithUser("Admin"), fabsdk.WithOrg("Org1"))
        chCtx, err := chProvider()

    2. 创建服务发现客户端discovery.Client请求，并添加PeersQuery；向某targetpeer 发送查询请求
        req := discclient.NewRequest().OfChannel(orgChannelID).AddConfigQuery()
        responses, err := client.Send(reqCtx, req,peerCfg1.PeerConfig)  

    3. 获取某channel中的所有peers，可以循环输出peers节点信息
        resp := responses[0]
        chanResp := resp.ForChannel(orgChannelID)
        configResult , err := chanResp.Config()
```

## 背书节点信息查询

> 背书节点的查询是跟背书策略有关系的
- 例如:
  - OR ('Org1MSP.peer','Org2MSP.peer')  
    > 背书节点只会查出来1个
  - AND ('Org1MSP.peer','Org2MSP.peer')  
    > 背书节点会查出来2个

- 获得背书节点信息内容，输出结果  

```bash

[
    {
        "Chaincode": "mycc",
        "EndorsersByGroups": {
            "G0": [
                {
                    "MSPID": "Org1MSP",
                    "LedgerHeight": 5,
                    "Endpoint": "peer0.org1.example.com:7051",
                    "Identity": "-----BEGIN CERTIFICATE-----\nMIICKDCCAc+gAwIBAgIRANTiKfUVHVGnrYVzEy1ZSKIwCgYIKoZIzj0EAwIwczEL\nMAkGA1UEBhMCVVMxEzARBgNVBAgTCkNhbGlmb3JuaWExFjAUBgNVBAcTDVNhbiBG\ncmFuY2lzY28xGTAXBgNVBAoTEG9yZzEuZXhhbXBsZS5jb20xHDAaBgNVBAMTE2Nh\nLm9yZzEuZXhhbXBsZS5jb20wHhcNMTgwNjA5MTE1ODI4WhcNMjgwNjA2MTE1ODI4\nWjBqMQswCQYDVQQGEwJVUzETMBEGA1UECBMKQ2FsaWZvcm5pYTEWMBQGA1UEBxMN\nU2FuIEZyYW5jaXNjbzENMAsGA1UECxMEcGVlcjEfMB0GA1UEAxMWcGVlcjAub3Jn\nMS5leGFtcGxlLmNvbTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABD8jGz1l5Rrw\n5UWqAYnc4JrR46mCYwHhHFgwydccuytb00ouD4rECiBsCaeZFr5tODAK70jFOP/k\n/CtORCDPQ02jTTBLMA4GA1UdDwEB/wQEAwIHgDAMBgNVHRMBAf8EAjAAMCsGA1Ud\nIwQkMCKAIOBdQLF+cMWa6e1p2CpOEx7SHUinzVvd55hLm7w6v72oMAoGCCqGSM49\nBAMCA0cAMEQCIC3bacbDYphXfHrNULxpV/zwD08t7hJxNe8MwgP8/48fAiBiC0cr\nu99oLsRNCFB7R3egyKg1YYao0KWTrr1T+rK9Bg==\n-----END CERTIFICATE-----\n"
                }
            ],
            "G1": [
                {
                    "MSPID": "Org2MSP",
                    "LedgerHeight": 5,
                    "Endpoint": "peer1.org2.example.com:7051",
                    "Identity": "-----BEGIN CERTIFICATE-----\nMIICKDCCAc+gAwIBAgIRAIs6fFxk4Y5cJxSwTjyJ9A8wCgYIKoZIzj0EAwIwczEL\nMAkGA1UEBhMCVVMxEzARBgNVBAgTCkNhbGlmb3JuaWExFjAUBgNVBAcTDVNhbiBG\ncmFuY2lzY28xGTAXBgNVBAoTEG9yZzIuZXhhbXBsZS5jb20xHDAaBgNVBAMTE2Nh\nLm9yZzIuZXhhbXBsZS5jb20wHhcNMTgwNjA5MTE1ODI4WhcNMjgwNjA2MTE1ODI4\nWjBqMQswCQYDVQQGEwJVUzETMBEGA1UECBMKQ2FsaWZvcm5pYTEWMBQGA1UEBxMN\nU2FuIEZyYW5jaXNjbzENMAsGA1UECxMEcGVlcjEfMB0GA1UEAxMWcGVlcjEub3Jn\nMi5leGFtcGxlLmNvbTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABOVFyWVmKZ25\nxDYV3xZBDX4gKQ7rAZfYgOu1djD9EHccZhJVPsdwSjbRsvrfs9Z8mMuwEeSWq/cq\n0cGrMKR93vKjTTBLMA4GA1UdDwEB/wQEAwIHgDAMBgNVHRMBAf8EAjAAMCsGA1Ud\nIwQkMCKAII5YgskKERCpC5MD7qBUQvSj7xFMgrb5zhCiHiSrE4KgMAoGCCqGSM49\nBAMCA0cAMEQCIDJmxseFul1GZ26djKa6jZ6zYYf6hchNF5xxMRWXpCnuAiBMf6JZ\njZjVM9F/OidQ2SBR7OZyMAzgXc5nAabWZpdkuQ==\n-----END CERTIFICATE-----\n"
                },
            ]
        },
        "Layouts": [
            {
                "quantities_by_group": {
                    "G0": 1,
                    "G1": 1
                }
            }
        ]
    }
]
```

- 代码示例  
> discovery client Request AddEndorsersQuery() --> Responses Get Endorsers()  

```bash
    1.  获取FabricSDK实例并获取ClientProvider
        sdk , err := fabsdk.New(configProvider,fabsdk.WithServicePkg(&DynamicDiscoveryProviderFactory{}))
        require.NoError(t, err, "Failed to create new SDK")
        defer sdk.Close()
        chProvider := sdk.Context(fabsdk.WithUser("Admin"), fabsdk.WithOrg("Org1"))
        chCtx, err := chProvider()

    2. 创建服务发现客户端discovery.Client请求，并添加PeersQuery；向某targetpeer 发送查询请求
        req := discclient.NewRequest().OfChannel(orgChannelID).AddEndorsersQuery()
        responses, err := client.Send(reqCtx, req, peerConfig)

    3. 获取某channel中的所有peers，可以循环输出peers节点信息
        resp := responses[0]
        chanResp := resp.ForChannel(orgChannelID)
        endorsers, err := chanResp.Endorsers(interest.Chaincodes, discclient.NewFilter(discclient.NoPriorities, discclient.NoExclusion))
```