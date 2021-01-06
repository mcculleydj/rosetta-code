<template>
  <v-app>
    <v-main>
      <v-container fluid>
        <v-row>
          <v-col cols="auto">
            <v-img src="logo.png" width="300" />
            <ProblemList @select="handleSelection($event)" />
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
              <p>
                I started this
                <a
                  href="http://www.rosettacode.org/wiki/Rosetta_Code"
                  target="_blank"
                  >Rosetta Code</a
                >
                clone to help prepare for upcoming technical interviews. It
                serves both as a refresher on basic algorithms and as an excuse
                to write some code in all three languages I am proficient with.
                I find that these exercises help me to think Go linguistically,
                Pythonically, or in the case of JS, just go with it. I've
                already been reminded of some of the nuances of each language
                within the first few implementations and I'm looking forward to
                the comparison for more complex problems in the near future.
              </p>
              <p>
                The prompts come from books I've purchased or the web. My
                general approach is to prototype the solution in Python and then
                deal with strong typing in the other two languages. Expect to
                see more comments in the Python for this reason. I will also add
                concurrency to some of the Go implementations where it makes
                sense in order to keep those patterns fresh in my mind. The
                problems are broken down by category in the list on the left.
                I'm using highlight.js to display the formatted code and Chris
                Hager's
                <a
                  href="https://www.metachris.com/2017/02/vuejs-syntax-highlighting-with-highlightjs/"
                  target="_blank"
                  >blog post</a
                >
                on incorporating it as a Vue directive.
              </p>
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

    handleSelection(problem) {
      this.problem = problem
      const el = document.getElementsByClassName('scroll')[0]
      el.scrollTop = 0
    },
  },
}
</script>

<style>
@import url('https://fonts.googleapis.com/css2?family=Fira+Code&display=swap');

/* vuetify is overruling highlightjs */
.theme--light.v-application code {
  background-color: #011627 !important;
  color: #d6deeb !important;
}

.hljs {
  font-family: 'Fira Code', monospace !important;
}
</style>

<style scoped>
.code {
  font-size: 18px;
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
