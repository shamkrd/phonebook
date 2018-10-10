class WebService {
    constructor() {
        this.BaseUrl = "http://localhost:5001/"
        this.Http = new XMLHttpRequest()
    }
    CreateRequest(method, params, type, contentType, callback) {
        let url = this.BaseUrl + method;
        let webService = this;
        webService.Http.open(type, url, true);
        webService.Http.setRequestHeader('Content-type', contentType);
     
        webService.Http.onreadystatechange = function () {
            if (webService.Http.readyState == 4 && webService.Http.status == 200) {
                callback();
            } 
        }
        webService.Http.send(JSON.stringify(params));
    }
}