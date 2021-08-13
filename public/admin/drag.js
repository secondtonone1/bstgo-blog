var dragSrcEl = null;

function dragStart(e) {
    dragSrcEl = this;
    e.dataTransfer.effectAllowed = 'move';
    e.dataTransfer.setData('text/html', this.outerHTML);

}

function dragOver(e) {
    if (e.preventDefault) {
        e.preventDefault();
    }

    if (dragSrcEl != null)
        this.classList.add('over');

    e.dataTransfer.dropEffect = 'move';

    return false;

}

function dragLeave(e) {
    this.classList.remove('over');

}

function drop(e) {
    if (e.stopPropagation) {
        e.stopPropagation();
    }

    if (dragSrcEl != this) {
        this.parentNode.removeChild(dragSrcEl);
        var dropHTML = e.dataTransfer.getData('text/html');
        this.insertAdjacentHTML('beforebegin', dropHTML);
        var dropElem = this.previousSibling;
        actions(dropElem);
    }
    this.classList.remove('over');
    dragSrcEl = null
        //发送保存目录排序请求
    let id = $(this).parent().parent().attr('id')
    console.log('id is ', id)
    let subCat = {}
    subCat.menu = []

    $(this).parent().children('.mini-li').each(function(index, element) {
        console.log(index)
        console.log(element)
        console.log($(element).attr('subcatid'))
        console.log($(element).children('span').text())
        let catele = {
            'catid': $(element).attr('subcatid'),
            'name': $(element).children('span').text(),
            'parent': id,
            'index': index
        }
        subCat.menu.push(catele)
    })
    console.log(subCat)
    console.log('JSON.stringify(subCat) is ', JSON.stringify(subCat))
    $.ajax({
        type: "POST",
        url: "/admin/sortmenu",
        contentType: "application/json",
        data: JSON.stringify(subCat), //参数列表
        dataType: "json",
        success: function(result) {
            //请求正确之后的操作
            console.log('post success , result is ', result)
        },
        error: function(XMLHttpRequest, textStatus, errorThrown) {
            //请求失败之后的操作
            console.log('post failed')
                // 状态码
            console.log(XMLHttpRequest.status);
            // 状态
            console.log(XMLHttpRequest.readyState);
            // 错误信息   
            console.log(textStatus);
        }
    });

    return false;

}

function dragEnd(e) {
    this.classList.remove('over');
    reorderItems();
    dragSrcEl = null
}

function reorderItems(e) {
    // var data = [];
    // for(var i = 0; i <= items.length; i++) {
    //   data.push({id: items[i].getAttribute('data-id'), order: i});
    // }
    // console.log(data);
}

function actions(elem) {
    console.log("action elem is ", elem)
    elem.addEventListener('dragstart', dragStart, false);
    elem.addEventListener('dragover', dragOver, false);
    elem.addEventListener('dragleave', dragLeave, false);
    elem.addEventListener('drop', drop, false);
    elem.addEventListener('dragend', dragEnd, false);
}

var items = document.getElementsByClassName('mini-li');
[].forEach.call(items, actions);