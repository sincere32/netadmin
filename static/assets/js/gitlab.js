function GetBranch(branch, repos_name) {
    branch.empty();
    $.ajax({
        type: "GET",
        url: "/v1.0.0/gitlab/repository/branches?repos_name=" + repos_name,
        success: function (data) {
            if (data.status != 0) {
                alert(data.msg);
            } else {
                branch.append("<option selected='selected'  value=-1>Select Branch</option>");
                for (var i=0;i<data.msg.length;i++){
                    branch.append("<option value=" + data['msg'][i]['name'] + ">" + data['msg'][i]['name'] + "</option>");
            }
            }
        }
    });
}

function GetRepository(repos_select) {
    $.ajax({
        type: "GET",
        url: "/v1.0.0/gitlab/repository",
        success: function (data) {
            if (data.status != 0) {
                alert("Get Repository Failed")
            } else {
                repos_select.empty();
                repos_select.append("<option selected='selected'  value=-1>Select Repository</option>");
                for (var i=0;i<data.data.length;i++){
                    repos_select.append("<option value=" + data['data'][i]['repos_name'] + ">" + data['data'][i]['repos_name'] + "</option>");
                }
            }

        }
    });
}

function GetBranchTree(repos_name, branch,file_select) {
    file_select.empty();
    $.ajax({
        type: "GET",
        url: "/v1.0.0/gitlab/repository/tree?repos_name=" + repos_name + "&branch=" + branch,
        success: function (data) {
            if (data.status != 0) {
                alert("Get Repository Failed")
            } else {
                file_select.append("<option selected='selected'  value=-1>default /</option>");
                for (var i = 0; i < data.data.length; i++) {
                    console.info(data.msg);
                    if (data.data[i]["type"] == "tree") {
                        file_select.append("<option value=" + i + ">" + data.data[i]["name"] + "</option>");
                    }
                }
            }
        }
    });
}

function GetFileList(repos_name, branch, file_select) {
    $.ajax({
        type: "GET",
        url: "/v1.0.0/gitlab/repository/tree?repos_name=" + repos_name + "&branch=" + branch,
        success: function (data) {
            if (data.status != 0) {
                alert("Get Repository Failed")
            } else {
                file_select.empty();
                file_select.append("<option selected='selected'  value=-1>Select File</option>");
                for (var i = 0; i < data.data.length; i++) {
                    if (data.data[i]["type"] == "blob") {
                        file_select.append("<option value=" + data.data[i]['id'] + ">" + data.data[i]["path"] + "</option>");
                    }
                }
            }

        }
    });
}


function GetFileContent(repos_name,blob_id, file_content) {
    $.ajax({
        type: "GET",
        url: "/v1.0.0/gitlab/repository/blobs?&repos_name=" + repos_name + "&blob_id=" + blob_id,
        success: function (data) {
            if (data.status != 0) {
                alert("Get Repository Failed")
            } else {
                file_content.val(data.msg);
            }
        }
    });
}