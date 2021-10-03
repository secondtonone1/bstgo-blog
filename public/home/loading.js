loadinst = {}

loadinst.loading = function() {
    console.log("loading....")
    $('.loadbkg').addClass('loading');
}

loadinst.unloading = function() {
    console.log("unloading ready")
    $('.loadbkg').removeClass('loading');
}