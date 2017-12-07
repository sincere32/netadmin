define(function(){

    return {
    'checkOnline' : function() {
        require(['jquery','jquery-cookie'], function (j, c) {

            j.ajax({
                type: "GET",
                url: "/v1.0.0/authentication/"+ c.cookie('nickname'),
                success: function (data) {
                    console.info(nickname);
                    console.info(data);
                    if (data.status != 0) {
                        window.location.href = "signin.html";
                    }
                }
            });

        });
    }
}
});


// (function(){
//
//     check_online();
//
// })();
// function check_online() {
//     console.info("123123123");
//     var nickname = getCookie("nickname");
//     $.ajax({
//         type: "GET",
//         url: "/v1.0.0/authentication/"+ nickname,
//         success: function (data) {
//             console.info(nickname);
//             console.info(data);
//             if (data.status != 0) {
//                 window.location.href = "signin.html";
//             }
//         }
//     });
// }
//
// function getCookie(name)
// {
//     var arr,reg=new RegExp("(^| )"+name+"=([^;]*)(;|$)");
//     if(arr=document.cookie.match(reg))
//         return unescape(arr[2]);
//     else
//         return null;
// }
//
