# 爬取n63明星海报图片
第一页:
明星列表页
找到
<table id="tablen63">
  <tbody>
    <tr>
      <!-- 这个是头 -->
      <td class="sb">
    </tr>
    <!-- 每一行数据 -->
    <tr>
      <td> 
        <a href="http://www.n63.com/n_china/2R">明星的名字</a>
        <a href="http://www.n63.com/n_china/2R/desktop">壁纸</a>
      </td>  
    </tr>
  </tbody>

遍历每个td item，获取href的链接，
继续访问这个链接
解析链接 / 分割字符串，得到最后一个字符串 2R 即为明星的名字

第二页：
具体某一个明星的图片列表,这一页呈现了明星的缩略图，存在分页的情况，需要获取页标签
需要获取全部页面总数
<div id="whole">
  <div id="container">
    <div id="imgholder">
      <a href="具体的地址"> //这个地址可能是另一个图片列表，也可能是原图的页面
    <div id="imgholder">
      <a href="具体的地址">
    <div id="imgholder">
      <a href="具体的地址">
点击缩略图的连接，进入实际的原图页面，这其中可能存在类似文件夹的情况

判断url如果包含字段 "www.n63.com/photodir/?album=" 说明是原图的链接，跳转到第三页，否则略过


第三页:
原图页面
<div id="image_content">
  <div id="image_show">
    <div id="imgholder">
      <a href="具体的地址"> 
        <img src="具体的图片地址">  //直接下载地址







http://www.n63.com/photodir/n63img/?N=X2hiJTI2ZGRXJTVEa1dmVyU1RFclNjBsVyU1RVclNURXZWclNUIlMjYlMkIuZiUyNm9vbyUyN2hnbGNrJTVEJTVDJTI3SiUyQSUyN1lmYSU2MCU1QiUyNw%3D%3D&v=.jpg


http://www.n63.com/photodir/img.php/thumbnail//china/2R/www.n63.com_e_f_th_n_t_z_ll.jpg



#tablen63 > tbody > tr:nth-child(2) > td:nth-child(2)
