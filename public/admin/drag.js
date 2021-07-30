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
    elem.addEventListener('dragstart', dragStart, false);
    elem.addEventListener('dragover', dragOver, false);
    elem.addEventListener('dragleave', dragLeave, false);
    elem.addEventListener('drop', drop, false);
    elem.addEventListener('dragend', dragEnd, false);
}

var items = document.getElementsByClassName('mini-li');
[].forEach.call(items, actions);