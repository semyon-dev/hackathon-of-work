<template>
    <div>

    <table>
        <tr v-for="(answer,question,index) in questionsAnswers" :key="index" >
            <td>
                {{question}}
            </td>
            <td>
                {{answer}}
            </td>
        </tr>
    </table>
</div>
</template>

<script>
    export default {
        name: "ResumeTable",
        props: ['QA'],
        data: ()=>{
            return{
                questionsAnswers: {}
            }
        },
        methods:{

        },
        created() {
            this.questionsAnswers = this.QA
        }
    }
</script>

<style scoped>

</style>
