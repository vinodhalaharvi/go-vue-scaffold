<template>
  <div>
    <form @submit.prevent="submitForm">
      <div v-for="(message, index) in formData.messages" :key="`message-${index}`" class="message-row">

        <!-- Role Icon Selectors -->
        <div class="role-icons">
          <i class="fas fa-server"
             :class="{ 'active': message.role === 'system' }"
             @click="setMessageRole(index, 'system')"></i>
          <i class="fas fa-user-secret"
             :class="{ 'active': message.role === 'assistant' }"
             @click="setMessageRole(index, 'assistant')"></i>
          <i class="fas fa-user"
             :class="{ 'active': message.role === 'user' }"
             @click="setMessageRole(index, 'user')"></i>
        </div>

        <!-- Content Input -->
        <input v-model="message.content" type="text" placeholder="Enter content" required>

        <!-- Remove Message Button -->
        <button @click="removeMessage(index)" type="button" class="remove-message-btn">
          <i class="fas fa-trash"></i>
        </button>
      </div>

      <button type="button" @click="addMessage" class="add-message-btn">
        <i class="fas fa-plus"></i> Add Message
      </button>
      <button type="submit" class="submit-btn">
        <i class="fas fa-paper-plane"></i> Send Request
      </button>
    </form>
  </div>
</template>

<script>
import '@fortawesome/fontawesome-free/css/all.css';


export default {
  data() {
    return {
      formData: {
        model: 'gpt-3.5-turbo',
        messages: [
          {
            role: 'system',
            content: '',
          },
        ],
      },
    };
  },
  methods: {
    setMessageRole(index, role) {
      this.formData.messages[index].role = role;
    },
    addMessage() {
      this.formData.messages.push({role: 'user', content: ''});
    },
    removeMessage(index) {
      this.formData.messages.splice(index, 1);
    },
    submitForm() {
      console.log('Submitting:', this.formData);
      // Here you would call your method to send the request to the API
    },
  },
};
</script>


<style scoped>
.message-row {
  display: flex;
  align-items: center;
  margin-bottom: 10px;
}

.role-icons i {
  cursor: pointer;
  margin: 0 10px;
  color: #ccc;
}

.role-icons i.active {
  color: #007bff;
}

input[type="text"] {
  flex-grow: 1;
  margin-right: 10px;
  padding: 8px;
  border: 1px solid #ccc;
  border-radius: 4px;
}

.remove-message-btn {
  background: none;
  border: none;
  cursor: pointer;
  color: #dc3545;
}

.remove-message-btn:hover {
  color: #c82333;
}

.add-message-btn, .submit-btn {
  margin-top: 20px;
  padding: 8px 16px;
  border: none;
  cursor: pointer;
  border-radius: 4px;
  color: white;
}

.add-message-btn {
  background-color: #28a745;
}

.add-message-btn:hover {
  background-color: #218838;
}

.submit-btn {
  background-color: #007bff;
}

.submit-btn:hover {
  background-color: #0056b3;
}
</style>