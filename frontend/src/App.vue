<template>
  <v-app>
    <v-main>
      <Problem v-if="problem" :problem="problem" @back="problem = null" />
      <ProblemList v-else @select="problem = $event" />
    </v-main>
  </v-app>
</template>

<script>
import ProblemList from '@/components/ProblemList'
import Problem from '@/components/Problem'

export default {
  components: {
    ProblemList,
    Problem,
  },

  data: () => ({
    problem: null,
  }),

  methods: {
    back() {
      this.problem = null
    },
  },
}
</script>

<style>
/* vuetify is overruling highlightjs */
.theme--light.v-application code {
  background-color: #011627 !important;
  color: #d6deeb !important;
}
</style>
