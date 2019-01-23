<template>
  <div class="hello">
    <ul>
      <li v-bind:key="c.name" v-for="c in companies">
        <b>{{ c.name }}:</b>
        {{ c.mission }}
      </li>
    </ul>
  </div>
</template>

<script>
export default {
  name: "Companies",

  computed: {
    // Return the list of companies we would get from the API
    companies() {
      return this.$store.companies;
    }
  },

  mounted() {
    // Refresh the list of companies
    this.$store.dispatch("refreshCompanies");
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
ul {
  list-style-type: none;
  padding: 0;
}
li {
  /*display: inline-block;*/
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
