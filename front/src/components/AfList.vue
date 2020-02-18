<template>
    <div class="ground">
      <div class="content">
        <div class="relate_wrap" id="divSearchRelate" v-if="related">
          <h4>연관검색어</h4>
          <ul>
            <li v-for="oData in relatedList" v-bind:key="oData">
              <a :href="searchUrl(oData.d)">{{oData.d}}</a>
            </li>
          </ul>
        </div>
        <router-view/>
      </div>
      <div class="aside">
        <!--연관검색 BJ리스트-->
        <div class="relate_bj" style="" id="relatebj_ul" v-if="tag">
          <h4 class="stit">연관검색 BJ 리스트</h4>
          <ul>
            <li v-for="oData in tagList" v-bind:key="oData">
              <a :href="searchUrl(oData.nick)">{{oData.nick}}</a>
            </li>
          </ul>
        </div>
        <!--//추천 검색어-->
        <div class="word" style="" id="recommend_Keyword" v-if="recommend">
          <h4 class="stit">추천 검색어</h4>
          <ul>
           <li v-for="oData in recommendList" v-bind:key="oData">
              <a :href="searchUrl(oData.keyword)">{{oData.keyword}}</a>
            </li>
          </ul>
        </div>
        <div class="live_word" style="" id="hot_Keyword">
          <h4 class="stit">실시간 검색어</h4>
          <ol>
            <li v-for="(oData, nIndex) in hotList" v-bind:key="oData">
              <a :href="searchUrl(oData.keyword)">
                <em>{{ (nIndex+1) }}</em>
                <span class="tit">{{ oData.keyword }}</span>
                <span class="rank">
                  <span :class="liveKeywordList(oData.updown, oData.show_text)">{{ oData.show_text }}</span>
                </span>
            </li>
          </ol>
        </div>
        <!--//#-->
      </div>
      <button type="button" class="btn_top" id="top_button" style="display: none;">top</button>
    </div>
</template>

<script>
export default {
  name: 'afList',
  data () {
    return {
      related: false,
      relatedList: '',
      tag: false,
      tagList: '',
      hot: true,
      hotList: '',
      recommend: true,
      recommendList: '',
      keyword: '',
      liveCnt: 0
    }
  },
  methods: {
    relatedKeyword () {
      let params = {}
      params['m'] = 'relatedSearch'
      params['v'] = '1.0'
      params['t'] = 'json'
      params['c'] = 'EUC-KR'
      params['d'] = this.keyword
      params['n'] = 20
      params['w'] = 'webk'

      this.$http.get('//scketc.afreecatv.com/api.php', {
        params: params
      }).then((res) => {
        if (res.data.list.length !== 0) {
          this.relatedList = res.data.list
          this.related = true
        } else {
          this.related = false
        }
      })
    },
    tagSearch () {
      let params = {}
      params['m'] = 'tagSearch'
      params['v'] = '1.0'
      params['t'] = 'json'
      params['c'] = 'EUC-KR'
      params['d'] = this.keyword
      params['n'] = 20
      params['w'] = 'webk'

      this.$http.get('//scketc.afreecatv.com/api.php', {
        params: params
      }).then((res) => {
        if (res.data.RELATED && res.data.RELATED.length !== 0) {
          this.tagList = res.data.RELATED
          this.tag = true
        } else {
          this.tag = false
        }
      })
    },
    recommendKeyword () {
      this.$http.get('//res.afreecatv.com/data/af_recommend_search_utf8.js', {
      }).then((res) => {
        let oResString = res.data.split('=')[1].replace(';', '')
        // oResString = iconv.encode(oResString, 'utf-8')
        let oResJson = JSON.parse(oResString)
        if (oResJson.LISTS[0].LIST.length !== 0) {
          this.recommendList = oResJson.LISTS[0].LIST
          this.recommend = true
        } else {
          this.recommend = false
        }
      })
    },
    hotKeyword () {
      let params = {}
      params['m'] = 'hotKeyword'
      params['v'] = '1.0'
      params['c'] = 'EUC-KR'

      this.$http.get('//scketc.afreecatv.com/api.php', {
        params: params
      }).then((res) => {
        if (res.data.HOT.length !== 0) {
          this.hotList = res.data.HOT
          this.hot = true
        }
      })
    },
    setting () {
      this.keyword = this.$route.query.keyword
    },
    searchUrl (szKeyword) {
      return '?keyword=' + encodeURIComponent(szKeyword)
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
    this.relatedKeyword()
    this.tagSearch()
    this.recommendKeyword()
    this.hotKeyword()
  }
}
</script>

<style></style>
