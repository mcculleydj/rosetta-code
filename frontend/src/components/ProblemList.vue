<template>
  <v-list class="pl-2">
    <v-list-group
      v-for="(category, i) in categories"
      :key="`category-${category}`"
      v-model="lists[i]"
    >
      <template v-slot:activator>
        <v-list-item-content>
          <v-list-item-title style="font-size: 20px">
            {{ category }}
          </v-list-item-title>
        </v-list-item-content>
      </template>

      <v-list-item
        v-for="problem in problems[category]"
        :key="`problem-${category}-${problem.number}`"
        :input-value="problem === selectedProblem"
        @click="onClick(problem)"
      >
        <v-list-item-icon>
          <v-icon>mdi-chevron-right</v-icon>
        </v-list-item-icon>
        <v-list-item-content>
          <v-list-item-title>
            {{ problem.title }}
          </v-list-item-title>
        </v-list-item-content>
      </v-list-item>
    </v-list-group>
  </v-list>
</template>

<script>
import problems from '@/../public/problems.json'

export default {
  props: ['backTrigger'],

  computed: {
    categories() {
      return Object.keys(problems)
    },
  },

  data: () => ({
    problems,
    selectedProblem: null,
    lists: [],
  }),

  methods: {
    onClick(problem) {
      this.selectedProblem = problem
      this.$emit('select', problem)
    },
  },

  watch: {
    backTrigger() {
      this.selectedProblem = null
      this.lists = this.lists.map(() => false)
    },
  },
}
</script>
