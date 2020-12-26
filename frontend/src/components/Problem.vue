<template>
  <v-container fluid>
    <v-row>
      <v-col>
        <v-btn color="primary" text @click="$emit('back')">
          <v-icon>mdi-chevron-left</v-icon>
          back to problem list
        </v-btn>
      </v-col>
    </v-row>
    <v-row>
      <v-col style="max-width: 700px;">
        <h2>{{ problem.title }}</h2>
        {{ problem.description }}
      </v-col>
    </v-row>
    <v-row>
      <v-col style="max-width: 700px;">
        <h3>Python</h3>
        <pre
          v-highlightjs="problem.python"
        ><code class="code python"></code></pre>
      </v-col>
      <v-col style="max-width: 700px;">
        <h3>Go</h3>
        <pre v-highlightjs="problem.go"><code class="code go"></code></pre>
      </v-col>
      <v-col style="max-width: 700px;">
        <h3>TypeScript</h3>
        <pre
          v-highlightjs="problem.ts"
        ><code class="code typescript"></code></pre>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
export default {
  props: ['problem'],
}
</script>

<style scoped>
.code {
  font-size: 16px;
  text-align: left;
  white-space: pre-wrap;
  word-wrap: normal;
  padding: 16px;
}
</style>
