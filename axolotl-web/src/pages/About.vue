<template>
  <component :is="$route.meta.layout || 'div'">
    <div class="about">
      <img class="logo" src="/axolotl.png" alt="Axolotl logo">
      <h1 class="title">Axolotl Beta {{ config.Version }}</h1>
      <h2 v-translate class="subtitle">A cross-platform Signal client</h2>
      <div class="description">
        <translate>
          This is a free and open source Signal client written in golang, rust and
          vuejs.
        </translate>
        <br>
        <translate class="mr-1">
          You can support the development of Axolotl either by filling
        </translate>
        <a
          v-translate
          href="https://github.com/nanu-c/axolotl/issues"
          @click="openExternal($event, 'https://github.com/nanu-c/axolotl/issues')"
        >issues at the bug tracker</a>.
        <br>
        <translate class="mr-1">or by becoming a</translate>
        <a
          href="https://www.patreon.com/bePatron?u=11219559"
          @click="
            openExternal($event, 'https://www.patreon.com/bePatron?u=11219559')
          "
        ><span v-translate>sponsor on patreon.</span></a><br>
        <br>
        <font-awesome-icon id="heart" icon="heart" />
        <br>
        <a
          href="https://axolotl.chat"
          @click="openExternal($event, 'https://axolotl.chat')"
        >https://axolotl.chat</a>
      </div>
    </div>
  </component>
</template>

<script>
import { mapState } from "vuex";
export default {
  name: "AboutPage",
  components: {},
  data() {
    return {
      showConfirmationModal: false,
    };
  },
  computed: mapState(["config", "gui"]),
  methods: {
    openExternal(e, url) {
      if (this.gui === "ut") {
        e.preventDefault();
        alert(url);
      }
    },
  },
};
</script>
<style scoped>
.about {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  text-align: center;
}
.logo {
  margin: 20px auto;
  border-radius: 10px;
}
#heart {
  font-size: 2rem;
  color: #0a6af4;
}
h1 {
  font-size: 1.5rem;
}
h2 {
  font-size: 1.3rem;
}
</style>

<!-- Add "scoped" attribute to limit CSS to this component only -->
