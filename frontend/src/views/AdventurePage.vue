<template>
  <div id="app">
    <h1>Solo Adventure Picker</h1>

    <div v-if="adventure" class="card">
      <h2>{{ displayName }}</h2>
      <h3>{{ displayType }}</h3>
      <!-- <p v-if="adventure.tags?.length">Tags: {{ adventure.tags.join(', ') }}</p> -->
    </div>

    <p v-if="errorMsg" class="error">{{ errorMsg }}</p>

    <button @click="getRandomAdventure">Pick Another Adventure</button>
  </div>
</template>

<script setup>
import { capitalizeWords } from "../utils/formatting.js";
import {ref, onMounted, computed} from 'vue'
import { useRoute } from 'vue-router'

const route = useRoute()
const adventure = ref(null)
const errorMsg = ref('')

const getRandomAdventure = async () => {
  const region = route.query.region || ''
  try {
    const res = await fetch(`http://localhost:8080/random?region=${region}`)

    if (!res.ok) {
      const errorJson = await res.json()
      throw new Error(errorJson.details || errorJson.error || 'Something went wrong')
    }

    const data = await res.json()
    adventure.value = data
    errorMsg.value = ''
  } catch (err) {
    adventure.value = null
    errorMsg.value = err.message
  }
}

onMounted(getRandomAdventure)

const displayName = computed(() => capitalizeWords(adventure.value?.name))
const displayType = computed(() => capitalizeWords(adventure.value?.type))
</script>

<style scoped>
.card h2, .card h3 {
  margin: 0.5rem 0;
  color: #213547;
}

button {
  margin-top: 2rem;
}
.error {
  color: red;
  margin-top: 1rem;
}
.card {
  background-color: #ffffff;
  border-radius: 12px;
  box-shadow: 0 4px 8px rgba(0,0,0,0.1);
  padding: 2rem;
  margin-top: 2rem;
}
</style>
