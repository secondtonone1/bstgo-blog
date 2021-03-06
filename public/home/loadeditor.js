function loadEditor() {
    window.editor = null
    const E = window.wangEditor
    const editor = new E('.comment-editor')
        //    设置编辑区域高度为 500 px
    editor.config.height = 100
        //配置菜单
    editor.config.menus = [
        'head',
        'bold',
        // 'fontSize',
        // 'fontName',
        'italic',
        //'underline',
        //'strikeThrough',
        // 'indent',
        //'lineHeight',
        //'foreColor',
        //'backColor',
        'link',
        'list',
        //'todo',
        // 'justify',
        'quote',
        'emoticon',
        //'image',
        //'video',
        //'table',
        'code',
        'splitLine',
        'undo',
        'redo',
    ]
    hljs.initHighlightingOnLoad(); // 初始化
    hljs.initLineNumbersOnLoad();
    editor.highlight = hljs
    editor.config.languageTab = '    '
    editor.config.pasteIgnoreImg = false
    editor.config.uploadImgShowBase64 = true
    editor.create()
    window.editor = editor
}