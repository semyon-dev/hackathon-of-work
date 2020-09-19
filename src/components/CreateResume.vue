<template>
  <div>
    <div class="right-cell-grid">
      <div class="center">
        <img alt="Vue logo" src="../assets/logo.jpg" width="260px">
        <div>
          <input list="competentions" type="text" v-model="compType"/>

          <datalist id="competentions">
            <option v-for="(comp, index) in competentions" :key="index">
              {{ comp }}
            </option>
          </datalist>

          <br>
          <br>

          <input  type="file"  v-on:change="changeFile" />

          <br>
          <br>


          <button class="white--text v-btn theme--light elevation-10 v-size--default accent" @click="load">загрузить резюме</button>
        </div>

      </div>

      <div>
        <div class="vector1">
          <img src="../images/Vector1.png" />
        </div>
        <div class="vector2">
          <img src="../images/Vector2.png" />
        </div>
      </div>


    </div>
  </div>
</template>

<script>



export default {
  name: "CreateResume",
  data: () => {
    return {
      compType: "",
      competentions: ["Официант", "Кладовищик", "Водитель погрузчика"],
      resume: ""
    }

  },
  methods: {
    load() {
      let url = ""
      let data = {context: this.resume, competention: this.compType}
      fetch(url, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
          // 'Content-Type': 'application/x-www-form-urlencoded',
        },
        body: JSON.stringify(data)

      }).then(response => response.json())
          .then(data => console.log(data));


    },
    changeFile(event){
      console.log('file changed!!!!!!!!!')
      console.log(event.target.files[0])

    }
  }
}
</script>

<style scoped src="../styles/allstyles.css">


  .center{
    alignment: center;
  }

</style>
