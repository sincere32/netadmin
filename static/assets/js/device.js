function GetHosts(host, type) {
    host.multiSelect();
    host.empty().multiSelect('refresh');
    $.ajax({
        type: "GET",
        url: "/v1.0.0/device/"+type,
        async: false,
        success: function (data) {
            for (var i = 0; i < data.data.length; i++) {
                $("#host").multiSelect('addOption', {value: data.data[i]['ip'], text: data.data[i]['ip'], index: i});
            }
        },
        error: function (data) {
            result.empty();
            result.append(data.msg);
        }
    });
}