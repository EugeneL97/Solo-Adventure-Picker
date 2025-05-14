<template>
  <div id="app">
    <h1>Solo Adventure Picker</h1>

    <Transition name="fade" mode="out-in">
      <div v-if="adventure" :key="adventure.name" class="card">
        <div class="xp-badge">+{{ adventure.xpValue || 100 }}XP</div>
        <div class="card-content">
          <h2>{{ displayName }}</h2>
          <h3>{{ displayType }}</h3>
          
          <div class="card-details" :class="{ expanded: isExpanded }">
            <p class="description">{{ adventure.description || 'Embark on this mysterious adventure!' }}</p>
            <div class="tags" v-if="adventure.tags?.length">
              <span class="tag" v-for="tag in adventure.tags" :key="tag">{{ tag }}</span>
            </div>
          </div>

          <button class="toggle-details" @click="isExpanded = !isExpanded">
            {{ isExpanded ? 'Show Less' : 'Show More' }}
          </button>
        </div>
      </div>
      <div v-else-if="errorMsg" class="error card">
        {{ errorMsg }}
      </div>
    </Transition>

    <div class="reroll-btn-wrapper">
      <button 
        class="reroll-btn" 
        @click="getRandomAdventure" 
        :disabled="isLoading || isInCooldown"
        :class="{ 'in-cooldown': isInCooldown }"
      >
        {{ buttonText }}
      </button>
    </div>
  </div>
</template>

<script setup>
import { capitalizeWords } from "../utils/formatting.js";
import {ref, onMounted, computed} from 'vue'
import { useRoute } from 'vue-router'

const route = useRoute()
const adventure = ref(null)
const errorMsg = ref('')
const isLoading = ref(false)
const isExpanded = ref(false)
const isInCooldown = ref(false)
const COOLDOWN_DURATION = 2000 // 2 seconds cooldown

const buttonText = computed(() => {
  if (isLoading.value) return 'Finding Adventure...'
  if (isInCooldown.value) return 'Wait a moment...'
  return 'Pick Another Adventure'
})

const startCooldown = () => {
  isInCooldown.value = true
  setTimeout(() => {
    isInCooldown.value = false
  }, COOLDOWN_DURATION)
}

const getRandomAdventure = async () => {
  if (isInCooldown.value || isLoading.value) return
  
  const region = route.query.region || ''
  isLoading.value = true
  isExpanded.value = false
  
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
  } finally {
    isLoading.value = false
    startCooldown()
  }
}

onMounted(getRandomAdventure)

const displayName = computed(() => capitalizeWords(adventure.value?.name))
const displayType = computed(() => capitalizeWords(adventure.value?.type))
</script>

<style scoped>
.card {
  background-color: #ffffff;
  border-radius: 12px;
  box-shadow: 0 4px 8px rgba(0,0,0,0.1);
  padding: 3.5rem 2rem 2rem 2rem;
  margin-top: 2rem;
  transition: all 0.3s ease;
  position: relative;
  max-width: 600px;
  margin-left: auto;
  margin-right: auto;
}

.card:hover {
  transform: translateY(-2px);
  box-shadow: 0 6px 12px rgba(0,0,0,0.15);
}

.card-content {
  text-align: center;
}

.xp-badge {
  position: absolute;
  top: 1rem;
  right: 1rem;
  background-color: #4CAF50;
  color: white;
  padding: 0.5rem 1rem;
  border-radius: 20px;
  font-weight: bold;
}

.card h2 {
  margin: 0 0 0.5rem 0;
  color: #213547;
  font-size: 1.8rem;
}

.card h3 {
  margin: 0 0 1.5rem 0;
  color: #4a5568;
  font-size: 1.2rem;
  font-weight: 500;
}

.card-details {
  max-height: 0;
  overflow: hidden;
  transition: max-height 0.3s ease-out;
  text-align: left;
}

.card-details.expanded {
  max-height: 500px;
}

.description {
  margin: 1rem 0;
  line-height: 1.6;
  color: #213547;
  text-align: center;
}

.tags {
  display: flex;
  flex-wrap: wrap;
  gap: 0.5rem;
  margin: 1rem 0;
  justify-content: center;
}

.tag {
  background-color: #e9ecef;
  padding: 0.3rem 0.8rem;
  border-radius: 15px;
  font-size: 0.9rem;
}

.toggle-details {
  background: none;
  border: none;
  color: #666;
  cursor: pointer;
  padding: 0.5rem;
  margin-top: 1rem;
  font-size: 0.9rem;
}

.toggle-details:hover {
  color: #333;
}

.reroll-btn {
  margin-top: 2rem;
  background-color: #646cff;
  color: white;
  border: none;
  padding: 0.8rem 1.6rem;
  border-radius: 8px;
  cursor: pointer;
  transform-origin: center;
  width: 220px;
  box-sizing: border-box;
  text-align: center;
  white-space: nowrap;
  overflow: hidden;
  max-width: 90vw;
  transition: background-color 0.3s ease;

}

.reroll-btn:hover:not(:disabled) {
  background-color: #747bff;
}

.reroll-btn:disabled {
  cursor: not-allowed;
}

.reroll-btn.in-cooldown {
  background-color: #8b8b8b;
  cursor: wait;
  animation: shrink-and-spring 2s cubic-bezier(0.22, 0.61, 0.36, 1);
}

@keyframes shrink-and-spring {
  0% {
    transform: scale(1);
    opacity: 1;
  }
  45% {
    transform: scale(0.2);
    opacity: 0.8;
  }
  50% {
    transform: scale(0.2);
    opacity: 0.8;
  }
  80% {
    transform: scale(2);
    opacity: 0.9;
  }
  100% {
    transform: scale(1);
    opacity: 1;
  }
}

.error {
  color: red;
  margin-top: 1rem;
}

/* Transition animations */
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.3s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>
