var aid = '0';
// var user = MyLocalStorage.get('user');
// if (user!=null) user = JSON.parse(user);
var user = "游客";
// if (user!=null) user = JSON.parse(user);
var limit = 3;
layui.use(['jquery', 'form', 'layedit','flow','util'], function () {
    var util = layui.util;
    var form = layui.form;
    var $ = layui.jquery;
    var layedit = layui.layedit;
    var flow = layui.flow;

    // 获取页面aid编号
    var url = window.location.href.split("?");
    if (url.length==2) {
        url = url[1];
        var args = url.split("&");
        aid = args[0].split("=")[1];
    }

    //评论和留言的编辑器
    var editIndex = layedit.build('remarkEditor', {
        height: 150,
        tool: ['face', '|','strong','italic', 'del','left', 'center', 'right', '|', 'link'],
    });

    //评论和留言的编辑器的验证
    layui.form.verify({
        content: function (value) {
            // if (user==null) return "登录后才能提交";
            value = $.trim(layedit.getText(editIndex));
            if (value == "") return "至少得有一个字吧";
            layedit.sync(editIndex);
        }
    });

    // 流加载 文章留言信息
    flow.load({
        elem: '.blog-comment', //流加载容器
        isAuto: false,
        done: function(page,next) {
            $.ajax({
                type: 'GET',
                data: {page:page,limit:limit,aid:aid},
                url: "/api/v1/getcomment",
                success:function(result) {
                    if (result.code==0) {
                        var msgs = result.data;
                        var lis = [];
                        for (var i=0; i<msgs.length; i++) {
                            var time = timestampToTime(msgs[i].create_time);
                            var html = '<li>'+
                                '<div class="comment-parent">'+
                                '<img src="static/frontend/images/user/15.jpg"/>'+
                                '<div class="info">'+
                                '<span class="username">游客</span>'+
                                '</div>'+
                                '<div class="content">'+
                                msgs[i].content+
                                '</div>'+
                                '<p class="info">'+
                                '<span class="time"><i class="fa fa-clock-o"></i>&nbsp;'+time+'</span>'+
                                '<span class="dh">'+
                                '<a class="btn-dzan" href="javascript:dzan(\''+msgs[i].id+'\');" id="dzan_'+msgs[i].id+'"><img src="static/frontend/images/zan.png"></img>'+msgs[i].like_count+'</a>'+
                                '<a class="btn-reply" href="javascript:btnReplyClick(\''+msgs[i].id+'\')" id="a_'+msgs[i].id+'");"><img src="static/frontend/images/huifu.png"></img>回复</a>'+
                                '</span>'+
                                '</p>'+
                                '</div>'+
                                '<hr />';
                            var cs = msgs[i].item;
                            if (cs!=null && cs.length>0) {
                                for (var j=0; j<cs.length; j++) {
                                    var time = timestampToTime(cs[j].create_time);
                                    html += '<div class="comment-child">'+
                                        '<img src="static/frontend/images/user/15.jpg" alt="Absolutely" />'+
                                        '<div class="info">'+
                                        '<span class="username">游客</span><span>'+cs[j].content+'</span>'+
                                        '</div>'+
                                        '<p class="info">'+
                                        '<span class="time"><i class="fa fa-clock-o"></i>&nbsp;'+time+'</span>'+
                                        '<span class="dh">'+
                                        '<a class="btn-dzan" href="javascript:dzan(\''+cs[j].id+'\');" id="dzan_'+cs[j].id+'"><img src="static/frontend/images/zan.png"></img>'+cs[j].like_count+'</a>'+
                                        '</span>'+
                                        '</p><hr/>'+
                                        '</div>';
                                }
                            }
                            html +='<div class="replycontainer layui-hide" id="'+msgs[i].id+'">'+
                                '<form class="layui-form" action="">'+
                                '<input type="hidden" name="reply_to" value="'+msgs[i].id+'">'+
                                '<div class="layui-form-item">'+
                                '<textarea name="replyContent" lay-verify="replyContent" id="t_'+msgs[i].id+'" placeholder="请输入回复内容" class="layui-textarea" style="min-height:80px;"></textarea>'+
                                '</div>'+
                                '<div class="layui-form-item">'+
                                '<button class="layui-btn layui-btn-mini" lay-submit="formReply" lay-filter="formReply">提交</button>'+
                                '</div>'+
                                '</form>'+
                                '</div></li>';
                            lis.push(html);
                        }
                        next(lis.join(''), page < Math.ceil(result.count/limit));
                    } else {
                        layer.msg(result.msg,{anim:6,icon:5});
                    }
                }
            });
        }
    });

    //监听留言提交
    form.on('submit(formLeaveMessage)', function (data) {
        var index = layer.load(1);
        //模拟留言提交
        setTimeout(function () {
            layer.close(index);
            var json = {aid:aid,content:data.field.editorContent,reply_to:0,uid:0};
            $.ajax({
                type: 'POST',
                data: json,
                url: "/api/v1/postcomment",
                success:function(result) {
                    if (result.code==0) {
                        var msg = result.data;
                        var time = timestampToTime(msg.CreateTime);
                        var html = '<li>'+
                            '<div class="comment-parent">'+
                            '<img src="static/frontend/images/user/15.jpg"/>'+
                            '<div class="info">'+
                            '<span class="username">游客</span>'+
                            '</div>'+
                            '<div class="content">'+
                            msg.Content+
                            '</div>'+
                            '<p class="info">'+
                            '<span class="time"><i class="fa fa-clock-o"></i>&nbsp;'+time+'</span>'+
                            '<span class="dh">'+
                            '<a class="btn-dzan" href="javascript:dzan(\''+msg.Id+'\');" id="dzan_'+msg.Id+'"><img src="static/frontend/images/zan.png"></img>'+msg.LikeCount+'</a>'+
                            '<a class="btn-reply" href="javascript:btnReplyClick(\''+msg.Id+'\')" id="a_'+msg.Id+'");"><img src="static/frontend/images/huifu.png"></img>回复</a>'+
                            '</span>'+
                            '</p>'+
                            '</div>'+
                            '<hr />'+
                            '<div class="replycontainer layui-hide" id="'+msg.Id+'">'+
                            '<form class="layui-form" action="">'+
                            '<input type="hidden" name="reply_to" value="'+msg.Id+'">'+
                            '<div class="layui-form-item">'+
                            '<textarea name="replyContent" lay-verify="replyContent" placeholder="请输入回复内容" class="layui-textarea" style="min-height:80px;"></textarea>'+
                            '</div>'+
                            '<div class="layui-form-item">'+
                            '<button class="layui-btn layui-btn-mini" lay-submit="formReply" lay-filter="formReply">提交</button>'+
                            '</div>'+
                            '</form>'+
                            '</div></li>';
                        $('.blog-comment').prepend(html);
                        $('#remarkEditor').val('');
                        editIndex = layui.layedit.build('remarkEditor', {
                            height: 150,
                            tool: ['face', '|', 'left', 'center', 'right', '|', 'link'],
                        });
                        layer.msg("留言成功", { icon: 1 });
                    } else {
                        layer.msg(result.message,{anim:6,icon:5});
                    }
                }
            });
        }, 500);
        return false;
    });

    //监听留言回复提交
    form.on('submit(formReply)', function (data) {
        // if (user==null) {
        //     layer.msg("登录后才能提交",{anim:6,icon:5});
        //     return false;
        // }
        if (data.field.replyContent == "") {
            layer.msg("至少得有一个字吧",{anim:6,icon:5});
            return false;
        }
        var index = layer.load(1);
        //模拟留言回复
        setTimeout(function () {
            layer.close(index);
            var json = {aid:aid,content:data.field.replyContent,reply_to:data.field.reply_to,uid:0};
            $.ajax({
                type: 'POST',
                data: json,
                url: "/api/v1/postcomment",
                success:function(result) {
                    if (result.code==0) {
                        var msg = result.data;
                        var time = timestampToTime(msg.CreateTime);
                        var html = '<div class="comment-child">'+
                            '<img src="static/frontend/images/user/15.jpg" alt="Absolutely" />'+
                            '<div class="info">'+
                            '<span class="username">游客</span><span>'+msg.Content+'</span>'+
                            '</div>'+
                            '<p class="info">'+
                            '<span class="time"><i class="fa fa-clock-o"></i>&nbsp;'+time+'</span>'+
                            '<span class="dh">'+
                            '<a class="btn-dzan" href="javascript:dzan(\''+msg.Id+'\');" id="dzan_'+msg.Id+'"><img src="static/frontend/images/zan.png"></img>'+msg.like_count+'</a>'+
                            '</span>'+
                            '</p><hr/>'+
                            '</div>';
                        $(data.form).find('textarea').val('');
                        $(data.form).parent('.replycontainer').before(html).siblings('.comment-parent').children('p').children('a').click();
                        layer.msg("回复成功", { icon: 1 });
                    } else {
                        layer.msg(result.message,{anim:6,icon:5});
                    }
                }
            });
        }, 500);
        return false;
    });
});

function btnReplyClick(elem) {
    var $ = layui.jquery;
    $('#'+elem).toggleClass('layui-hide');
    if ($('#a_'+elem).text() == '回复') {
        $('#a_'+elem).html('<i class="fa fa-caret-square-o-up" style="font-size:18px;"></i>&nbsp;收起');
    } else {
        $('#a_'+elem).html('<img src="images/huifu.png"></img>回复');
    };
}

function dzan(mid) {
    layer.msg("待开发");
    // var i = parseInt($('#dzan_'+mid).text());
    // i++;
    // $.ajax({
    //     type: 'POST',
    //     data: {dzan:i,mid:mid},
    //     url: "",
    //     success:function(result) {
    //         if (result.code==1) {
    //             $('#dzan_'+mid).html('<img src="static/frontend/images/zan_d.png" class="animated bounceIn"></img>'+i);
    //         } else {
    //             layer.msg(result.msg,{anim:6,icon:5});
    //         }
    //     }
    // });

}