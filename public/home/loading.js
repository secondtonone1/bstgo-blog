var demo = [];
demo.loading = function() { //定义显示loading的js方法
    console.log('demo.loading....')
        //$("#loading").css("left", ($(window).width() - $("#loading").outerWidth()) / 2 + $(window).scrollLeft() + "px");
    $("#whiteOverlay").css("height", document.body.clientHeight + "px");
    $("#whiteOverlay").css("width", document.body.clientWidth + "px");
    //$("#loading").css("top", ($(window).height() - $("#loading").outerHeight()) / 2 + $(window).scrollTop() + "px");
    $("#loading").show();
    $("#whiteOverlay").show();
};

demo.hiding = function() { //定义隐藏loading的就是方法
    console.log('demo.hiding....')
    $("#loading").hide();
    $("#whiteOverlay").hide()
};

$(document).ready(function() {
    //  demo.loading(); //页面加载的时候调用显示loading框。
});

$(window).load(function() {
    demo.hiding()
});