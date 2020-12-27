<template>
  <v-app>
    <v-main>
      <v-container fluid>
        <v-row>
          <v-col cols="auto">
            <v-img src="logo.png" width="300" />
            <ProblemList @select="problem = $event" />
          </v-col>
          <v-col class="scroll mt-5">
            <template v-if="problem">
              <h2 class="mb-3">{{ problem.title }}</h2>
              <span class="desc-txt">{{ problem.description }}</span>
              <v-row justify="center" class="mt-3">
                <v-col cols="auto">
                  <v-img src="python.png" width="50" class="lang-img" />
                </v-col>
              </v-row>
              <pre
                v-highlightjs="problem.python"
              ><code class="code python"></code></pre>

              <v-row justify="center" class="mt-3">
                <v-col cols="auto">
                  <v-img src="go.png" width="50" class="lang-img" />
                </v-col>
              </v-row>
              <pre
                v-highlightjs="problem.go"
              ><code class="code go"></code></pre>

              <v-row justify="center" class="mt-3">
                <v-col cols="auto">
                  <v-img src="ts.png" width="50" class="lang-img" />
                </v-col>
              </v-row>
              <pre
                v-highlightjs="problem.ts"
              ><code class="code typescript"></code></pre>
            </template>
            <div v-else class="desc-txt">
              A series of HackerRank-like problems solved in the three languages
              I know best. I find that these exercises help me to think Go
              linguistically, Pythonically, or whatever the hell the JS/TS
              version of that is.
            </div>
          </v-col>
        </v-row>
      </v-container>
    </v-main>
  </v-app>
</template>

<script>
import ProblemList from '@/components/ProblemList'

export default {
  components: {
    ProblemList,
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

<style scoped>
.code {
  font-size: 20px;
  text-align: left;
  white-space: pre-wrap;
  word-wrap: normal;
  padding: 32px 16px 16px 16px;
}

.scroll {
  overflow-y: scroll;
  max-height: 98vh;
}

.lang-img {
  margin-bottom: -25px;
}

.desc-txt {
  font-size: 20px;
  color: rgba(0, 0, 0, 0.57);
}
</style>
