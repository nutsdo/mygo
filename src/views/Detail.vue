<template>
  <div class="home">
    <div class="album  bg-light">
      <a :href="product.item_url"><img :src="product.small_images.string[1]" class="rounded mx-auto d-block" width="100%" :alt="product.title"></a>

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
    name : "detail",
    data :function() {
        return {
            product:{}
        }
    },
    mounted : function () {
        this.axios
            .get("http://localhost:8090/product/detail/" + this.$route.params.productId)
            .then(res=> {
                console.log(this)
                console.log(res.data.tbk_item_info_get_response.results.n_tbk_item)
                this.product = res.data.tbk_item_info_get_response.results.n_tbk_item[0]
            })
    },

}
</script>
