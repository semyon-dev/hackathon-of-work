<template>
  <div>

    <div class="right-cell-grid">
      <div class="center">
        <img alt="Vue logo" src="../assets/logo.jpg" width="260px">
        <div>
          <input list="competentions" type="text" v-model="compType"/>

          <datalist id="competentions">
            <option v-for="(comp, index) in competentions" :key="index">
              {{ comp.split('.')[0] }}
            </option>
          </datalist>

          <br>
          <br>

          <textarea v-model="resume" ></textarea>

          <br>
          <br>

          <button @click="load">загрузить резюме</button>
        </div>
        <div>


          <div class="vector2">
            <img src="../images/Vector2.png" />
          </div>

        </div>

      </div>
    </div>

    <div v-if="state===2">
      <h3>
        Resume Table:
      </h3>
      <div>
        <table>
          <tr v-for="(answer,question,index) in QA" :key="index" >
            <td>
              {{question}}
            </td>
            <td>
              {{answer.split(',')[0].slice(2)}}
            </td>
          </tr>
        </table>
      </div>


      <button @click="showRawJson">raw json</button>
      <div>
        {{rawJson}}
      </div>



    </div>


  </div>
</template>

<script>



export default {
  name: "CreateResume",
  data: () => {
    return {
      state: 1,
      compType: "",
      competentions: ["Официант", "Кладовищик", "Водитель погрузчика"],
      resume: "",
      QA: {},
      rawJson:"",
      models: []
    }

  },
  methods: {
    load() {
      let url = 'http://localhost:5000/resume'
      let data = {context: this.resume, competention: this.compType}
      fetch(url, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)

      }).then(response => response.json())
              .then(data => {
                console.log(data)
                this.QA = JSON.parse(data.data)
                console.log(this.QA)
                this.state=2
              });


    },
    changeFile(event){
      console.log('file changed!!!!!!!!!')
      console.log(event.target.files[0])

    },
    showRawJson(){
      let model = this.QA
      let newQA = {}
      let keys = Object.keys(model)
      keys.map((val)=>{
        newQA[val] = model[val].split(',')[0].slice(2)
      })

      this.rawJson = JSON.stringify(newQA)
    },
    getAvailableModels(){
      let url = 'http://localhost:5000/modelsAvailable'
      fetch(url, {
        method: 'GET',
      }).then(response => response.json())
              .then(data => {
                console.log(data)
                this.models = JSON.parse(data.models)
                console.log(this.models)
                this.competentions = this.models
              });
    }
  },
  created() {
    console.log('CREATED!!!!')
    this.getAvailableModels()
  }
}
</script>

<style scoped src="../styles/allstyles.css">


  .center{
    alignment: center;
  }

</style>
