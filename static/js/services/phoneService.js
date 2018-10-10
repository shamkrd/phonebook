class PhoneService {
    constructor(webService) {
        this.WebService = webService
    }


    add(phone) {
        let ws = new WebService();
        ws.CreateRequest("addPhone", phone, "POST", "application/json", function () {
            showToast(ws.Http.responseText);
        });
    }

    delete(phone) {
        let ws = new WebService();
        ws.CreateRequest("deletePhone", phone, "POST", "application/json", function () {
            showToast(ws.Http.responseText);
        });
    }

    update(phone) {
        let ws = new WebService();
        ws.CreateRequest("updatePhone", phone, "POST", "application/json", function () {
            showToast(ws.Http.responseText);
        });
    }
}