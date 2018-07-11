<template>
  <div class="videohosting">
    <h1>{{ title }}</h1>

    <h3 v-if="error">Ошибка: {{error}}</h3>

     <table>
      <tr>
        <th>Название</th>
        <th>Описание</th>
        <th>Ссылка на видео</th>
      </tr>
      <tr v-for="video in video_list" v-bind:key="video.title">
        <td>{{video.title}}</td>
        <td>{{video.description}}</td>
        <td>{{video.link}}</td>
        <td><button v-on:click="edit_video(video)">Редактировать</button></td>
        <td><button v-on:click="remove_video(video)">Удалить видеофайл</button> </td>
      </tr>
    </table>

     <h3 v-if="edit_index == -1">Новый видеофайл</h3>
     <form>
       <p>Название:<input type="text" v-model="new_video.title"></p>
       <p>Описание:<input type="text" v-model="new_video.description"></p>
       <p>Ссылка на видео:<input type="text" v-model="new_video.link"></p>
       <button v-if="edit_index == -1" v-on:click="add_new_video">Добавить видеофайл</button>
       <button v-if="edit_index > -1" v-on:click="end_of_edition">Сохранить изменения</button>
     </form>
  </div>
</template>

<script>
const axios = require('axios')

export default {
  name: 'videohosting',
  props: {
    title: String
  },
  data: function () {
    return {
      edit_index: -1,
      error: '',
      video_list: [],
      new_video:
       {
         'title': '',
         'description': '',
         'link': ''
       }
    }
  },
  mounted: function () {
    this.get_videos()
  },
  methods: {
    get_videos: function () {
      this.error = ''
      const url = '/api/videohosting/videos'
      axios.get(url).then(response => {
        this.video_list = response.data
      }).catch(response => {
        this.error = response.response.data
      })
    },
    add_new_video: function () {
      this.error = ''
      const url = '/api/videohosting/videos'
      axios.post(url, this.new_video).then(response => {
        this.video_list.push(this.new_video)
      }).catch(response => {
        this.error = response.response.data
      })
    },
    remove_video: function (item) {
      this.error = ''
      const url = '/api/videohosting/videos' + this.video_list.indexOf(item)
      axios.delete(url).then(response => {
        this.video_list.splice(this.video_list.indexOf(item), 1)
      }).catch(response => {
        this.error = response.response.data
      })
    },
    edit_video: function (item) {
      this.edit_index = this.video_list.indexOf(item)
      this.new_video = this.video_list[this.edit_index]
    },
    end_of_edition: function () {
      this.error = ''
      const url = '/api/videohosting/videos' + this.edit_index
      axios.put(url, this.new_video).then(response => {
        this.edit_index = -1
        this.new_video = {
          'title': '',
          'description': '',
          'link': ''
        }
      }).catch(response => {
        this.error = response.response.data
      })
    }
  }
}
</script>
<style scoped>
</style>
