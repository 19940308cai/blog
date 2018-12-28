/**
 * 本文件用于全选checkbox框的选择操作
 * 
 * 需要：
 * 		全选input框增加class名：t-selectAll
 * 		成员input框增加class名：t-checkbox
 * 
 */

// 反选
// $(function(){

// 	// 点击全选时
// 	$(".t-selectAll").click(function(event){
// 		$(".t-checkbox").each(function(){
// 			this.checked = !this.checked;
// 		});
// 	})
// });
// 
// $(function() {

    //全选/全不选
    // $(".t-selectAll").click(function() {
    //     $(".t-checkbox").prop("checked", this.checked);
    // });
    // $(".t-checkbox").click(function() {
    //     var flag = true;
    //     $(".t-checkbox").each(function() {
    //         if (!this.checked) {
    //             flag = false;
    //         }
    //     });
    //     $(".t-selectAll").prop("checked", flag);
    // })
// })

/**
 *  @from peng
 */

function allCheck(settings) {
    var options = {
        single: 't-checkbox',
        allCheck: 't-selectAll'
    }

    $.extend(options, settings);

    var $single = $('.' + options.single),
        $allCheck = $('.' + options.allCheck);

    $allCheck.change(function() {
        $single.prop('checked', this.checked);
    })

    $single.change(function() {
        $single.length === $('.' + options.single + ':checked').length ? $allCheck.prop('checked', true) : $allCheck.prop('checked', false)
    })
}
