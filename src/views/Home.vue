<template>
  <div class="home">
    <div class="album py-5 bg-light">
      <div class="container">

        <div class="row">
          <div class="col-md-3" v-for="product in products">
            <div class="card mb-3 box-shadow">
              <img class="card-img-top" :data-src="product.pict_url" alt="Thumbnail [100%x225]" style="height: 255px; width: 100%; display: block;" :src="product.pict_url" data-holder-rendered="true">
              <div class="card-body">
                <p class="card-text">{{ product.title }}</p>
                <div class="d-flex justify-content-between align-items-center">

                  <small class="text-muted">销量:{{ product.volume }}件</small>
                  <div class="btn-group">
                    <!--<router-link :to="{ name: 'detail', params: { productId: 123 }}" type="button" class="btn btn-sm btn-outline-secondary">User</router-link>-->
                    <button type="button" class="btn btn-sm btn-outline-secondary">¥{{ product.zk_final_price }}</button>
                    <button v-on:click="viewProduct" :data-productId="product.num_iid" type="button" class="btn btn-sm btn-outline-secondary">购买</button>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

  </div>
</template>

<script>
// @ is an alias to /src
// import HelloWorld from '@/components/HelloWorld.vue'
//
// export default {
//   name: 'home',
//   components: {
//     HelloWorld
//   }
// }

export default {
    name : "home",
    data :function() {
        return {
            products : []
        }
    },
    beforeCreate : function () {

        this.axios
            .get("http://localhost:8090/taobao/couponBrandRecommendSc")
            .then(res=> {
                // console.log(this)
                console.log(res.data.tbk_sc_coupon_brand_recommend_response.results.tbk_coupon)
                this.products = res.data.tbk_sc_coupon_brand_recommend_response.results.tbk_coupon
            })
    },
    methods: {
        viewProduct: function (event) {
            // `this` 在方法里指向当前 Vue 实例
            // `event` 是原生 DOM 事件
            if (event) {
                var proId = event.target.getAttribute('data-productId')
                console.log(proId)
                console.log(event.target.tagName)
                // window.location.href=""
                this.$router.push({name:"product.detail",params:{productId:proId}})
            }
        }
    }
}
</script>
