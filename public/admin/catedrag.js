var dragSrcEl_article = null;

function dragStart_article(e) {
    console.log('dragStart_article ', e)
    dragSrcEl_article = this;
    e.dataTransfer.effectAllowed = 'move';
    e.dataTransfer.setData('text/html', this.outerHTML);
}

function dragOver_article(e) {
    if (e.preventDefault) {
        e.preventDefault();
    }
    this.classList.add('over');

    e.dataTransfer.dropEffect = 'move';

    return false;

}

function dragLeave_article(e) {
    this.classList.remove('over');
}

function drop_article(e) {
    if (e.stopPropagation) {
        e.stopPropagation();
    }

    if (dragSrcEl_article != this) {
        this.parentNode.removeChild(dragSrcEl_article);
        var dropHTML = e.dataTransfer.getData('text/html');
        this.insertAdjacentHTML('beforebegin', dropHTML);
        var dropElem = this.previousSibling;
        actions_article(dropElem);
    }

    this.classList.remove('over');
    return false;

}

function dragEnd_article(e) {
    this.classList.remove('over');
    reorderItems_article();

}

function reorderItems_article(e) {
    // var data = [];
    // for(var i = 0; i <= items.length; i++) {
    //   data.push({id: items[i].getAttribute('data-id'), order: i});
    // }
    // console.log(data);
}

function actions_article(elem) {
    console.log('actions_article is ', elem)
    elem.addEventListener('selectstart', (e) => { return false }, true)
    elem.addEventListener('dragstart', dragStart_article, true);
    elem.addEventListener('dragover', dragOver_article, true);
    elem.addEventListener('dragleave', dragLeave_article, true);
    elem.addEventListener('drop', drop_article, true);
    elem.addEventListener('dragend', dragEnd_article, true);
}