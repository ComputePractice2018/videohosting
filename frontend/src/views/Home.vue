<template>
  <div class="home">
    <videohosting title="Видеохостинг"/>
  </div>
</template>

<script>
// @ is an alias to /src
import videohosting from '@/components/videohosting.vue'

export default {
  name: 'home',
  components: {
    videohosting
  }
}
</script>
