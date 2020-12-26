<template>
  <v-container fluid>
    <h2>Problems Grouped by Category</h2>
    <v-row>
      <v-col style="max-width: 800px">
        <v-list>
          <v-list-group
            v-for="category in categories"
            :key="`category-${category}`"
          >
            <template v-slot:activator>
              <v-list-item-content>
                <v-list-item-title>
                  {{ category }}
                </v-list-item-title>
              </v-list-item-content>
            </template>

            <v-list-item
              v-for="problem in problems[category]"
              :key="`problem-${category}-${problem.number}`"
              @click="$emit('select', problem)"
            >
              <v-list-item-content>
                <v-list-item-title>
                  {{ problem.title }}
                </v-list-item-title>
              </v-list-item-content>
            </v-list-item>
          </v-list-group>
        </v-list>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import problems from '@/../public/problems.json'

export default {
  computed: {
    categories() {
      return Object.keys(problems)
    },
  },

  data: () => ({
    problems,
  }),
}
</script>

<style scoped></style>
