function openAddContactModal() {
    document.getElementById("modal-title").textContent = "Добавление нового контакта";
    var nameInput = `<div class="row">
                        <div class="input-field col s6">
                            <input placeholder="Петров Петр" maxlength="200" id="contact-name" type="text" class="validate">
                            <label class="active" for="first_name">Имя контакта</label>
                        </div>
                    </div>`;
    document.getElementsByTagName("form")[0].innerHTML = nameInput;
    document.getElementById("modal-footer").firstElementChild.setAttribute("onclick", "addContact()");
    document.getElementById("modal-footer").firstElementChild.textContent = "Добавить";
}

function openPhonesModal(element) {
    var contact = new Contact();
    contact.ID = parseInt(element.getAttribute("data-id"));
    var phones;
    var ws = new WebService();
    ws.CreateRequest("getPhones", contact, "POST", "application/json", function (phones) {
        phones = JSON.parse(ws.Http.responseText).Phones;
        document.getElementById("modal-title").textContent = "Список телефонов";
        document.getElementsByTagName("form")[0].innerHTML = getPhonesTable(phones, contact.ID);
        document.getElementById("modal-footer").firstElementChild.textContent = "";
    });
}

function openAddPhoneModal(contactId) {
  document.getElementById("modal-title").textContent =
    "Добавление нового телефона";
  var nameInput = `<div class="row">
                        <div class="input-field col s6">
                            <input placeholder="79585354777" maxlength="11" id="contact-phone" type="text" class="validate">
                             <input id="contact-id" value="${contactId}" type="hidden" class="validate">
                            <label class="active" for="first_name">Номер телефона</label>
                        </div>
                    </div>`;
  document.getElementsByTagName("form")[0].innerHTML = nameInput;
  document
    .getElementById("modal-footer")
        .firstElementChild.setAttribute("onclick", "addPhone()");
    document.getElementById("modal-footer").firstElementChild.textContent = "Добавить";
}



function openUpdateContactModal(element) {
    var id = element.getAttribute("data-id");
    var name = element.getAttribute("data-name");
    document.getElementById("modal-title").textContent = "Редактирование контакта";
    var nameInput = `<div class="row">
                        <div class="input-field col s6">
                            <input value="${name}" maxlength="200" id="contact-name" type="text" class="validate">
                            <input type="hidden" id="contact-id" value="${id}">
                            <label class="active" for="first_name">Имя контакта</label>
                        </div>
                    </div>`;
    document.getElementsByTagName("form")[0].innerHTML = nameInput;
    document.getElementById("modal-footer").firstElementChild.setAttribute(`onclick`, `updateContact(${id})`);
    document.getElementById("modal-footer").firstElementChild.textContent = "Сохранить";
}

function openUpdatePhoneModal(element) {
  var id = element.getAttribute("data-id");
  var phone = element.getAttribute("data-name");
  document.getElementById("modal-title").textContent =
    "Редактирование контакта";
  var nameInput = `<div class="row">
                        <div class="input-field col s6">
                            <input value="${phone}" maxlength="11" id="contact-phone" type="text" class="validate">
                            <input type="hidden" id="phone-id" value="${id}">
                            <label class="active" for="first_name">Номер телефона</label>
                        </div>
                    </div>`;
  document.getElementsByTagName("form")[0].innerHTML = nameInput;
  document
    .getElementById("modal-footer")
    .firstElementChild.setAttribute(`onclick`, `updatePhone(${id})`);
  document.getElementById("modal-footer").firstElementChild.textContent =
    "Сохранить";
}

function openDeletePhoneModal(element) {
  var id = element.getAttribute("data-id");
  var name = element.getAttribute("data-name");
  document.getElementById("modal-title").textContent = "Удаление номера телефона";
  var nameInput = `<div class="row">
                       <h5>Вы действительно хотите удалить номер: <b>${name}</b> ?</h5>
                    </div>`;
  document.getElementsByTagName("form")[0].innerHTML = nameInput;
  document.getElementById("modal-footer")
    .firstElementChild.setAttribute("onclick", `deletePhone(${id})`);
  document.getElementById("modal-footer").firstElementChild.textContent =
    "удалить";
}


function openDeleteContactModal(element) {
    var id = element.getAttribute("data-id");
    var name = element.getAttribute("data-name");
    document.getElementById("modal-title").textContent = "Удаление контакта";
    var nameInput = `<div class="row">
                       <h5>Вы действительно хотите удалить контак: <b>${name}</b> ?</h5>
                    </div>`;
    document.getElementsByTagName("form")[0].innerHTML = nameInput;
    document.getElementById("modal-footer").firstElementChild.setAttribute("onclick", `deleteContact(${id})`);
    document.getElementById("modal-footer").firstElementChild.textContent = "удалить";
}