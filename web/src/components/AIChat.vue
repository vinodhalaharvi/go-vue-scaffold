<template>
  <div>
    <form @submit.prevent="submitForm">
      <div>
        <label for="model">Model:</label>
        <input id="model" v-model="formData.model" type="text" placeholder="Enter model" required>
      </div>

      <div v-for="(message, index) in formData.messages" :key="index">
        <label for="role">Role:</label>
        <select v-model="message.role" required>
          <option value="system">System</option>
          <option value="user">User</option>
          <option value="assistant">Assistant</option>
        </select>

        <label for="content">Content:</label>
        <input v-model="message.content" type="text" placeholder="Enter content" required>
      </div>

      <div>
        <button class="remove-message-btn" type="button" @click="removeMessage(index)">
          <i class="fas fa-trash-alt"></i>
        </button>
        <button type="button" @click="addMessage">
          <i class="fas fa-plus"></i>
        </button>
        <button type="submit">
          <i class="fas fa-paper-plane"></i>
        </button>

      </div>
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
form {
  max-width: 600px;
  margin: 20px auto;
  padding: 20px;
  border-radius: 8px;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
}

h3 {
  margin-bottom: 10px;
}

label {
  display: block;
  margin-bottom: 5px;
}

input[type="text"],
select {
  width: 100%;
  padding: 10px;
  margin-bottom: 20px;
  border-radius: 4px;
  border: 1px solid #ccc;
}

button {
  cursor: pointer;
  padding: 10px 20px;
  border: none;
  border-radius: 4px;
  background-color: #007bff;
  color: white;
  margin-right: 10px;
}

button:hover {
  background-color: #0056b3;
}

button[type="submit"] {
  background-color: #28a745;
}

button[type="submit"]:hover {
  background-color: #218838;
}

button i {
  margin-right: 5px;
}

.remove-message-btn {
  background-color: #dc3545;
}

.remove-message-btn:hover {
  background-color: #c82333;
}
</style>
