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
              <p>
                I started this Rosetta Code clone to help prepare for upcoming
                technical interviews. It serves both as a refresher on basic
                algorithms and as an excuse to write some code in all three
                languages I am proficient with. I find that these exercises help
                me to think Go linguistically, Pythonically, or whatever the
                hell the JS/TS version of that is. I've already been reminded of
                some of the nuances of each language within the first few
                implementations and I'm looking forward to the comparison across
                languages for more complex problems in the near future.
              </p>
              <p>
                The problems come from books I've purchased or the web. My
                approach is to prototype the solution in Python and then add
                strong typing by reimplementing the solution in TS and Go.
                Expect to see more comments in the Python for this reason. I
                will also add concurrency to the Go implementations for more
                complex algorithms just to keep those patterns fresh in my mind.
                The problems are broken down by category in the list on the
                left. I'm using highlight.js to display the formatted code and
                Chris Hager's
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
