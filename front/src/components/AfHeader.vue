<template>
  <div class="gnb_wrap">
    <div class="search_area">
      <h3>
        <a href=".">검색</a>
      </h3>
      <fieldset>
        <legend>검색 폼</legend>
        <input type="text" v-model="keyword" name="szKeyword" id="szKeyword" title="검색어 입력" class="search_bar" @keyup="autoJaso" @click=checkLive(keyword)>
        <a href="javascript:;" title="검색" class="btn_search" @click="search(keyword)">검색</a>
        <!-- 자동 완성 -->
        <div class="sear_auto" id="divSearchAuto" v-if="auto">
          <ul>
            <li v-for="oData in list" v-bind:key="oData">
              <a :href="searchUrl(oData.d)">{{oData.d}}</a>
            </li>
          </ul>
        </div>
        <!-- 최근 검색어 -->
        <div class="mysch" v-if="mysch">
          {{ list }}
        </div>
        <!-- 실시간 검색어 -->
        <div class="livesch" v-if="live">
          <p class="stitle">실시간 인기 검색어</p>
          <ul  v-for="(oData, nIndex) in list" v-bind:key="oData">
            <li>
              <a :href="searchUrl(oData.keyword)">
              <em>{{ (nIndex+1) }}</em>
              <span class="tit">{{ oData.keyword }}</span>
              <span class="rank">
                <span :class="liveKeywordList(oData.updown, oData.show_text)">{{ oData.show_text }}</span>

              </span>
              </a>
            </li>
          </ul>
        </div>
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
      list: '',
      mysch: false,
      auto: false,
      live: false,
      keyword: '',
      ticket: ''
    }
  },
  methods: {
    autoJaso (e) {
      let szKeyword = e.target.value
      if (!szKeyword) {
        this.checkLive()
        return false
      }
      // 방향키
      var aDirectionCode = [37, 38, 30, 40]
      if (aDirectionCode.includes(e.which)) {
      } else if (e.which === 13) {
        location.href = '//search.afreecatv.com?keyword=' + e.target.value
      } else {
        let params = {}
        params['m'] = 'autoJaso'
        params['v'] = '1.0'
        params['t'] = 'json'
        params['c'] = 'UTF-8'
        params['d'] = e.target.value
        params['n'] = 10
        params['w'] = 'webk'

        this.$http.get('//scketc.afreecatv.com/api.php', {
          params: params
        }).then((res) => {
          this.list = res.data.list
          this.live = false
          this.auto = true
        })
      }
    },
    search (e) {
      location.href = '//search.afreecatv.com?keyword=' + e
    },
    setting () {
      this.keyword = this.$route.query.keyword
    },
    searchUrl (szKeyword) {
      return '?keyword=' + szKeyword
    },
    checkLive (szKeyword) {
      if (szKeyword) {

      } else {
        let params = {}
        params['m'] = 'hotKeyword'
        params['v'] = '1.0'
        params['t'] = 'json'
        params['c'] = 'UTF-8'

        this.$http.get('//scketc.afreecatv.com/api.php', {
          params: params
        }).then((res) => {
          this.list = res.data.HOT
          this.auto = false
          this.live = true
        })
      }
    },
    liveKeywordList (szUpdown, szShowText) {
      if (szUpdown === 'up' || szUpdown === 'down') {
        return szUpdown
      } else {
        if (szShowText === 'new') {
          return szShowText
        } else {
          return szShowText
        }
      }
    }

  },
  created: function () {
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
