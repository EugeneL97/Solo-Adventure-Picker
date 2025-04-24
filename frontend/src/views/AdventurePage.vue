<template>
  <div id="app">
    <h1>Solo Adventure Picker</h1>

    <div v-if="adventure" class="card">
      <h2>{{ displayName }}</h2>
      <h2>{{ displayType }}</h2>
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
.error {
  color: red;
  margin-top: 1rem;
}
.card {
  border: 1px solid #ccc;
  padding: 1rem;
  margin-top: 1rem;
}
</style>
