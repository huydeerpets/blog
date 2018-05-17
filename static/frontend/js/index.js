layui.use(['jquery','flow','carousel'], function () {
    var $ = layui.jquery;

    var width = width || window.innerWeight || document.documentElement.clientWidth || document.body.clientWidth;
    width = width>1200 ? 1170 : (width > 992 ? 962 : width);
    var carousel = layui.carousel;
    //建造实例
    carousel.render({
        elem: '#carousel'
        ,width: width+'px' //设置容器宽度
        ,height:'360px'
        ,indicator: 'inside'
        ,arrow: 'always' //始终显示箭头
        ,anim: 'default' //切换动画方式

    });

    // 流加载 文章列表信息
    var inittag = $('#hiddentag').val();
    var flow = layui.flow;
    var limit = 7;
    flow.load({
        elem: '.blog-main-left', //流加载容器
        isAuto: false,
        end: '没有更多的文章了~QAQ',
        done: function(page,next) {
            $.ajax({
                type: 'GET',
                data: {page:page,limit:limit, tag:inittag},
                url: 'ajaxarticletable',
                success:function(result) {
                    if (result.code==0) {
                        var arts = result.data;
                        var lis = [];
                        for (var i=0; i<arts.length; i++) {
                            var head='';
                            if (page==1) {
                                head='<div class="article shadow animated fadeInLeft">';
                            } else {
                                head='<div class="article shadow animated zoomIn">';
                            }

                            head += '<div class="article-left ">'+
                                '<img src="'+arts[i].thumb+'" alt=\''+arts[i].title +'\'/>'+
                                '</div>'+
                                '<div class="article-right">'+
                                '<div class="article-title">'+
                                '<a href="fdetail?id='+arts[i].id+'">'+arts[i].title+'</a>'+
                                '</div>'+
                                '<div class="article-abstract">'+
                                arts[i].summary+
                                '</div>'+
                                '</div>'+
                                '<div class="clear"></div>'+
                                '<div class="article-footer">'+
                                '<span><i class="fa fa-clock-o"></i>&nbsp;&nbsp;'+timestampToTime(arts[i].create_time)+'</span>'+
                                '<span class="article-author"><i class="fa fa-user"></i>&nbsp;&nbsp;'+arts[i].author_name+'</span>';
                            if(arts[i].tags.length>0){
                                for(var j=0; j<arts[i].tags.length; j++){
                                    head += '<span><i class="fa fa-tag"></i>&nbsp;&nbsp;<a href="#">'+arts[i].tags[j].name+'</a></span>';
                                }
                            }

                            head +=  '<span class="article-viewinfo"><i class="fa fa-eye"></i>&nbsp;'+arts[i].san_count+'</span>'+
                                '<span class="article-viewinfo"><i class="fa fa-commenting"></i>&nbsp;'+arts[i].comment_count+'</span>'+
                                '</div>'+
                                '</div>';

                            lis.push(head);
                        }

                        if (page==1) {
                            next(lis.join(''), page < Math.ceil(result['count']/limit));
                        } else {
                            setTimeout(function(){next(lis.join(''), page < Math.ceil(result['count']/limit));},1000);
                        }
                    } else {
                        layer.msg(result.msg,{anim:6});
                    }
                }
            });
        }
    });

    $(function () {
        //播放公告
        playAnnouncement(5000);
    });
    function playAnnouncement(interval) {
        var index = 0;
        var $announcement = $('.home-tips-container>span');
        //自动轮换
        setInterval(function () {
            index++;    //下标更新
            if (index >= $announcement.length) {
                index = 0;
            }
            $announcement.eq(index).stop(true, true).fadeIn().siblings('span').fadeOut();  //下标对应的图片显示，同辈元素隐藏
        }, interval);
    }
});