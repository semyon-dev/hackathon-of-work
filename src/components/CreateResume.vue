<template>
  <div>
    <div>
      <input list="competentions" type="text" v-model="compType"/>

      <datalist id="competentions">
        <option v-for="(comp, index) in competentions" :key="index">
          {{ comp }}
        </option>
      </datalist>

      <br>
      <br>

      <textarea v-model="resume" style="align-content: center"></textarea>

      <br>
      <br>


      <button @click="load">send resume</button>

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


    }
  }
}
</script>

<style scoped>

</style>