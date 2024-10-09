<template>
  <div v-if="questions.length > 0">
    <div v-for="question in questions" :key="question.question">
      <p>{{ question.question }}</p>
      <el-radio-group v-model="answers[question.question]">
        <el-radio v-if="question.options && question.options.A" :label="question.options.A">{{ question.options.A }}</el-radio>
        <el-radio v-if="question.options && question.options.B" :label="question.options.B">{{ question.options.B }}</el-radio>
      </el-radio-group>
    </div>
    <el-button @click="submit">提交表单</el-button>
    <p>{{analyze}}</p>
  </div>
  <div v-else>
    <p>正在加载数据...</p>
  </div>
</template>

<script>
import axios from "axios";

export default {
  data() {
    return {
      questions: [],
      answers: {},
      analyze:"",
    };
  },
  mounted() {
    this.btnclicked();
  },
  methods: {
    cleanMarkdown(text) {
      // 如果这里的所谓 Markdown 标记只是特定字符，可以直接替换
      return text.replace(/```json/g, '').replace(/```/g, '');
    },
    submit() {

      axios.post('http://localhost:8080/submitAnswers', this.answers)
          .then(response => {
            console.log(response.data.message);
            this.analyze = response.data.message;
          })
          .catch(error => {
            console.error(error);
          });
    },
    btnclicked: function () {
      axios.get('http://localhost:8080/getMbit')
          .then(response => {
            // 处理请求成功的结果


            this.questions = JSON.parse('{' + this.cleanMarkdown(response.data.question)+ '}').questions;
            console.log(this.questions.questions);
          })
          .catch(error => {
            // 处理请求失败的情况
            console.error(error);
          });
    }
  }
};
</script>