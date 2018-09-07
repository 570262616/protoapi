/**
* This file is generated by 'protoapi'
* The file contains frontend API code that work with the library 'VueResource', therefore, it's required that 'VueResource' is installed in the project
* The generated code is written in TypeScript
* The code provides a basic usage for API call and may need adjustment according to specific project requirement and situation
* -------------------------------------------
* 该文件生成于protoapi
* 文件包含前端调用API的代码，并使用第三方库VueResource， 因此需要保证VueResource存在于项目中
* 文件内代码使用TypeScript
* 该生成文件只提供前端API调用基本代码，实际情况可能需要根据具体项目具体要求不同而作出更改
*/
import { Component, Vue } from 'vue-property-decorator';
import VueResource from 'vue-resource';
import {
    EnvListRequest, 
    EnvListResponse, 
    Error, 
    KVHistoryRequest, 
    KVHistoryResponse, 
    KeyListRequest, 
    KeyListResponse, 
    KeyValueListRequest, 
    KeyValueListResponse, 
    KeyValueRequest, 
    KeyValueResponse, 
    ProductListRequest, 
    ProductListResponse, 
    RegisterServiceRequest, 
    RegisterServiceResponse, 
    SearchKeyValueListRequest, 
    ServiceListRequest, 
    ServiceListResponse, 
    ServiceSearchRequest, 
    TagListRequest, 
    TagListResponse, 
    UpdateServiceRequest, 
    UpdateServiceResponse, 
    UploadProtoFileRequest, 
    UploadProtoFileResponse, 
    
} from './data';
import { generateUrl, errorHandling } from './helper';

Vue.use(VueResource);

@Component
export default class AppService extends Vue {
    // constructor
    constructor() {
        super()
    }

    // Base Url
    private baseUrl: string = "http://localhost:8080"
    getEnv(params: EnvListRequest): PromiseLike<EnvListResponse | Error> {

        let url: string = generateUrl(this.baseUrl,"AppService", "getEnv");

        return this.$http.post(url,params).then(
            res => {
                // handle success data - 200
                return Promise.resolve(res.data as EnvListResponse)
            },
            err => { 
                // handle error response
                return errorHandling(err.response)
            }
        );
    }
    
    registerService(params: RegisterServiceRequest): PromiseLike<RegisterServiceResponse | Error> {

        let url: string = generateUrl(this.baseUrl,"AppService", "registerService");

        return this.$http.post(url,params).then(
            res => {
                // handle success data - 200
                return Promise.resolve(res.data as RegisterServiceResponse)
            },
            err => { 
                // handle error response
                return errorHandling(err.response)
            }
        );
    }
    
    updateService(params: UpdateServiceRequest): PromiseLike<UpdateServiceResponse | Error> {

        let url: string = generateUrl(this.baseUrl,"AppService", "updateService");

        return this.$http.post(url,params).then(
            res => {
                // handle success data - 200
                return Promise.resolve(res.data as UpdateServiceResponse)
            },
            err => { 
                // handle error response
                return errorHandling(err.response)
            }
        );
    }
    
    uploadProtoFile(params: UploadProtoFileRequest): PromiseLike<UploadProtoFileResponse | Error> {

        let url: string = generateUrl(this.baseUrl,"AppService", "uploadProtoFile");

        return this.$http.post(url,params).then(
            res => {
                // handle success data - 200
                return Promise.resolve(res.data as UploadProtoFileResponse)
            },
            err => { 
                // handle error response
                return errorHandling(err.response)
            }
        );
    }
    
    getTags(params: TagListRequest): PromiseLike<TagListResponse | Error> {

        let url: string = generateUrl(this.baseUrl,"AppService", "getTags");

        return this.$http.post(url,params).then(
            res => {
                // handle success data - 200
                return Promise.resolve(res.data as TagListResponse)
            },
            err => { 
                // handle error response
                return errorHandling(err.response)
            }
        );
    }
    
    getProducts(params: ProductListRequest): PromiseLike<ProductListResponse | Error> {

        let url: string = generateUrl(this.baseUrl,"AppService", "getProducts");

        return this.$http.post(url,params).then(
            res => {
                // handle success data - 200
                return Promise.resolve(res.data as ProductListResponse)
            },
            err => { 
                // handle error response
                return errorHandling(err.response)
            }
        );
    }
    
    getServices(params: ServiceListRequest): PromiseLike<ServiceListResponse | Error> {

        let url: string = generateUrl(this.baseUrl,"AppService", "getServices");

        return this.$http.post(url,params).then(
            res => {
                // handle success data - 200
                return Promise.resolve(res.data as ServiceListResponse)
            },
            err => { 
                // handle error response
                return errorHandling(err.response)
            }
        );
    }
    
    searchServices(params: ServiceSearchRequest): PromiseLike<ServiceListResponse | Error> {

        let url: string = generateUrl(this.baseUrl,"AppService", "searchServices");

        return this.$http.post(url,params).then(
            res => {
                // handle success data - 200
                return Promise.resolve(res.data as ServiceListResponse)
            },
            err => { 
                // handle error response
                return errorHandling(err.response)
            }
        );
    }
    
    getKeyList(params: KeyListRequest): PromiseLike<KeyListResponse | Error> {

        let url: string = generateUrl(this.baseUrl,"AppService", "getKeyList");

        return this.$http.post(url,params).then(
            res => {
                // handle success data - 200
                return Promise.resolve(res.data as KeyListResponse)
            },
            err => { 
                // handle error response
                return errorHandling(err.response)
            }
        );
    }
    
    getKeyValueList(params: KeyValueListRequest): PromiseLike<KeyValueListResponse | Error> {

        let url: string = generateUrl(this.baseUrl,"AppService", "getKeyValueList");

        return this.$http.post(url,params).then(
            res => {
                // handle success data - 200
                return Promise.resolve(res.data as KeyValueListResponse)
            },
            err => { 
                // handle error response
                return errorHandling(err.response)
            }
        );
    }
    
    searchKeyValueList(params: SearchKeyValueListRequest): PromiseLike<KeyValueListResponse | Error> {

        let url: string = generateUrl(this.baseUrl,"AppService", "searchKeyValueList");

        return this.$http.post(url,params).then(
            res => {
                // handle success data - 200
                return Promise.resolve(res.data as KeyValueListResponse)
            },
            err => { 
                // handle error response
                return errorHandling(err.response)
            }
        );
    }
    
    updateKeyValue(params: KeyValueRequest): PromiseLike<KeyValueResponse | Error> {

        let url: string = generateUrl(this.baseUrl,"AppService", "updateKeyValue");

        return this.$http.post(url,params).then(
            res => {
                // handle success data - 200
                return Promise.resolve(res.data as KeyValueResponse)
            },
            err => { 
                // handle error response
                return errorHandling(err.response)
            }
        );
    }
    
    fetchKeyHistory(params: KVHistoryRequest): PromiseLike<KVHistoryResponse | Error> {

        let url: string = generateUrl(this.baseUrl,"AppService", "fetchKeyHistory");

        return this.$http.post(url,params).then(
            res => {
                // handle success data - 200
                return Promise.resolve(res.data as KVHistoryResponse)
            },
            err => { 
                // handle error response
                return errorHandling(err.response)
            }
        );
    }
    
}