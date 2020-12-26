<template>
  <div id="app">
    <h1>hello world</h1>
    <pre
      v-highlightjs
      class="code"
    ><code class="typescript">{{ test }}</code></pre>
  </div>
</template>

<script>
import test from '../public/iterables-1.ts.txt'
export default {
  data: () => ({ test }),
}
</script>

<style scoped>
.code {
  font-size: 16px;
  font-family: monospace;
  text-align: left;
}
</style>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
