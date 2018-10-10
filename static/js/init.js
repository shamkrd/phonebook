document.addEventListener('DOMContentLoaded', function () {
    var elems = document.querySelectorAll('.modal');
    var instances = M.Modal.init(elems, null);
});

let webService = new WebService();
let contactService = new ContactService(webService);
let phoneService = new PhoneService(webService)
let contact;
let phone;



function addPhone() {
    phone = new Phone(0, document.getElementById("contact-phone").value, parseInt(document.getElementById("contact-id").value));
    phoneService.add(phone);
}

function updatePhone(id) {
    phone = new Phone(id, document.getElementById("contact-phone").value, 0);
    phoneService.update(phone);
}

function deletePhone(id) {
    phone = new Phone(id, "", 0);
    phoneService.delete(phone);
}


function addContact() {
    contact = new Contact(0, document.getElementById("contact-name").value);
    contactService.add(contact);
}

function deleteContact(id) {
    contact = new Contact(id, "");
    contactService.delete(contact);
}

function updateContact(id) {
    contact = new Contact(id, document.getElementById("contact-name").value);
    contactService.update(contact);
}