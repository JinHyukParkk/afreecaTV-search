<template>
  <div id="header">
    <div id="gArea">
      <h1>
        <a href="//www.afreecatv.com" target="_top">afreecaTV</a>
      </h1>
    </div>
    <div class="gnb_wrap">
      <div class="search_area">
        <h3>
          <a href=".">검색</a>
        </h3>
        <fieldset>
          <legend>검색 폼</legend>
          <input type="text" v-model="keyword" name="szKeyword" id="szKeyword" title="검색어 입력" class="search_bar" @keyup="autoJaso">
          <a href="javascript:;" title="검색" class="btn_search" @click="search">검색</a>
          <!-- 자동 완성 -->
          <div class="sear_auto" id="divSearchAuto" v-if="auto">
            <p>몰라이씨</p>
          </div>
          <!-- 최근 검색어 -->
          <div class="mysch" v-if="mysch">
            {{ list }}
          </div>
          <!-- 실시간 검색어 -->
          <div class="livesch" style="display:none;"></div>
        </fieldset>
      </div>
  </div>
</template>

<script>
export default {
  name: 'afHeader',
  data () {
    return {
      msg: '재원공주다',
      list: '123',
      mysch: false,
      auto: true,
      keyword: ''
    }
  },
  methods: {
    autoJaso (e) {
      // 방향키
      var aDirectionCode = [37, 38, 30, 40]
      if (aDirectionCode.includes(e.which)) {
        console.log('key')
      } else if (e.which === 13) {
        location.href = '//search.afreecatv.com?keyword=' + e.target.value
      } else {
        this.$http.get('http://203.238.139.11:8080/test').then((res) => {
          console.log('result:', res.data)
          this.list = res.data
          this.mysch = true
        })
      }
    },
    search () {
      location.href = '//search.afreecatv.com?keyword=' + 'value'
    },
    setting() {
      console.log('이건 머야')
    }
  },
  created: function() {
    this.setting()
  }
}
</script>

<style>
#header {
  position: relative;
  width: 960px;
  margin: 0 auto;
  z-index: 1000;
}
#gArea a {
    text-decoration: none;
    display: block;
    height: 50px;
    text-indent: -9999px;
}

#gArea {
  position: relative;
  top: 0;
  left: 0;
  right: 0;
  z-index: 1500;
  width: 100%;
  min-width: 960px;
  height: 34px;
}
#header h1 {
  display: block;
  overflow: hidden;
  position: absolute;
  top: 0;
  left: 0;
  background: url(//www.afreecatv.com/images/svg/logo.svg) 0 50% no-repeat;
  width: 96px;
  height: 50px;
  margin: 0;
}

</style>
