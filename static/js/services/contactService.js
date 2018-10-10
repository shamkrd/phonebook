class ContactService{
    constructor(webService){
        this.WebService = webService
    }
   

    add(contact) {
        let webService = this.WebService;  
        webService.CreateRequest("createContact", contact, "POST", "application/json", function () {
            var customResponse = JSON.parse(webService.Http.responseText);
            var tr = createTr(customResponse.ID, contact.Name);
            appendElement(tr, "contacts");
            showToast(webService.Http.responseText);
        });
    }

    delete(contact) {
        let webService = this.WebService;
        webService.CreateRequest("deleteContact", contact, "POST", "application/json", function () {
            showToast(webService.Http.responseText);
            removeTr(contact.ID);
        });
    }

    update(contact) {
        let webService = this.WebService;
        webService.CreateRequest("updateContact", contact, "POST", "application/json", function () {
            showToast(webService.Http.responseText);
            let tr = document.getElementById(`tr-${contact.ID}`);
            tr.getElementsByClassName("name")[0].textContent = contact.Name;
            tr = createTr(contact.ID, contact.Name);
        });
    }
}