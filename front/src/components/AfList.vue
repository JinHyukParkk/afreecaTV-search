<template>
    <div id="ground">
      <div id="content">
        <div class="relate_wrap" id="divSearchRelate" v-if="related">
          <h4>연관검색어</h4>
          <ul>
            <li v-for="oData in relatedList" v-bind:key="oData">
              <a :href="searchUrl(oData.d)">{{oData.d}}</a>
            </li>
          </ul>
        </div>
      </div>
      <div id="aside"></div>
      <button type="button" class="btn_top" id="top_button" style="display: none;">top</button>
    </div>
</template>

<script>
export default {
  name: 'afList',
  data () {
    return {
      msg: '재원공주다',
      related: false,
      relatedList: '',
      keyword: '',
      ticket: ''
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
          console.log(res.data.list)
          this.relatedList = res.data.list
          this.related = true
        }
      })
    },
    setting () {
      this.keyword = this.$route.query.keyword
    },
    searchUrl (szKeyword) {
      return '?keyword=' + szKeyword
    }
  },
  created: function () {
    this.setting()
    this.relatedKeyword()
  }
}
</script>

<style></style>
