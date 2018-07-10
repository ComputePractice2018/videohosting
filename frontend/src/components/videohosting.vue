<template>
  <div class="videohosting">
    <h1>{{ title }}</h1>

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
       <button v-if="edit_index > -1" v-on:click="make_new_video">Сохранить изменения</button>
     </form>
  </div>
</template>

<script>
export default {
  name: 'videohosting',
  props: {
    title: String
  },
  data: function () {
    return {
      edit_index: -1,
      video_list: [
        {
          'title': 'little cat',
          'description': 'Это видео о котятах',
          'link': 'Ссылка на видео'
        },
        {
          'title': 'Hазвание',
          'description': 'Описание',
          'link': 'Ссылка на видео'
        }
      ],
      new_video:
       {
         'title': '',
         'description': '',
         'link': ''
       }
    }
  },
  methods: {
    add_new_video: function () {
      this.video_list.push(this.new_video)
    },
    remove_video: function (item) {
      this.video_list.splice(this.video_list.indexOf(item), 1)
    },
    edit_video: function (item) {
      this.edit_index = this.video_list.indexOf(item)
      this.new_video = this.video_list[this.edit_index]
    },
    make_new_video: function () {
      this.edit_index = -1
      this.new_video = {
        'title': '',
        'description': '',
        'link': ''
      }
    }
  }
}
</script>
<style scoped>
</style>
