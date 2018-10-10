function showToast(text) {
    M.toast({ html: text, options: { displayLength: 200 } });
}

function appendElement(element, parentId) {
    var parent =  document.getElementById(parentId);
    parent.insertAdjacentHTML('beforeend', element);
}

function removeTr(id) {
    var tr = document.getElementById(`tr-${id}`);
    tr.remove();
}

function getPhonesTable(phones, contactId) {
    let response = `<div class="row">
                      <table class="highlight">
                    <thead>
                        <th>ID</th>
                        <th>Номер телефона</th>
                        <th><a href="#modal" class="waves-effect modal-trigger" onclick="openAddPhoneModal(${contactId})"
                                    title="добавить номер телефона">
                                    <i class="btn material-icons">add</i>
                                </a></th>
                    </thead>
                    <tbody id="phones">
                 `;
    for (let i = 0; i < phones.length; i++) {
        response += `     
            <tr style="vertical-align: middle;" id="ptr-${phones[i].ID}">
                            <td>${phones[i].ID}</td>
                            <td class="name">${phones[i].Phone}</td>
                            <td>
                                <a href="#modal" data-id="${phones[i].ID}" data-name="${phones[i].Phone}" class="waves-effect modal-trigger" onclick="openUpdatePhoneModal(this)"
                                    title="Редактировать">
                                    <i class="btn material-icons">edit</i>
                                </a>
                                <a href="#modal" data-id="${phones[i].ID}" data-name="${phones[i].Phone}" class="waves-effect modal-trigger" onclick="openDeletePhoneModal(this)"
                                    title="Удалить">
                                    <i class="btn material-icons">delete</i>
                                </a>
                            </td>
                        </tr>`
    }
    response += `</tbody><table></div>`;
    return response;
}

function createTr(id, name) {
    var tr = `<tr style="vertical-align: middle;" id="tr-${id}">
                    <td>${id}</td>
                    <td class="name">${name}</td>
                    <td>
                         <a href="#modal" data-id="${id}" onclick="openPhonesModal(this)" title="Список телефонов"
                                    class="waves-effect modal-trigger">
                                    <i class="btn material-icons">local_phone</i>
                                </a>
                        <a href="#modal" data-id="${id}"
                        data-name="${name}" class="waves-effect modal-trigger" onclick="openUpdateContactModal(this)"  title="Редактировать">
                            <i class="btn material-icons">edit</i>
                        </a>
                        <a href="#modal" data-id="${id}" data-name="${name}" 
                        class="waves-effect modal-trigger"
                        onclick="openDeleteContactModal(this)" title="Удалить">
                            <i class="btn material-icons">delete</i>
                        </a>
                    </td>
                </tr>`;

    return tr;
}