import { useEffect, useState } from "react";
import { getTodos, addTodo, deleteTodo, toggleTodo } from "../api/todo";
import { Sparkles, Plus, Trash2, Check, Calendar, Zap, LogOut, Loader2, MoveRight, X } from "lucide-react";
import { useNavigate } from "react-router-dom";

interface Todo {
  id: number;
  task: string;
  completed: boolean;
}

export default function Todos() {
  const [todos, setTodos] = useState<Todo[]>([]);
  const [task, setTask] = useState("");
  const [aiPrompt, setAiPrompt] = useState("");
  const [aiSuggestions, setAiSuggestions] = useState<string[]>([]);
  const [loadingAI, setLoadingAI] = useState(false);
  const navigate = useNavigate();

  const fetchTodos = async () => {
    try {
      const res = await getTodos();
      setTodos(res.data);
    } catch (error) {
      console.error("Failed to fetch todos:", error);
    }
  };

  useEffect(() => {
    fetchTodos();
  }, []);

  const handleAdd = async () => {
    if (!task.trim()) return;
    try {
      await addTodo(task);
      setTask("");
      fetchTodos();
    } catch (error) {
      alert("Failed to add task");
    }
  };

  const handleDelete = async (id: number) => {
    try {
      await deleteTodo(id);
      fetchTodos();
    } catch (error) {
      alert("Failed to delete task");
    }
  };

  const handleToggle = async (id: number) => {
    try {
      await toggleTodo(id);
      fetchTodos();
    } catch (error) {
      alert("Failed to update task");
    }
  };

  const handleGenerateAITasks = async () => {
    if (!aiPrompt.trim()) return;
    setLoadingAI(true);
    try {
      const aiBaseUrl = import.meta.env.VITE_AI_API || "http://localhost:8082";
      // Call AI service to get suggestions
      const token = localStorage.getItem("token") || "";
      const response = await fetch(`${aiBaseUrl}/generate`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify({ prompt: aiPrompt }),
      });
      
      const data = await response.json();
      const tasks = data.tasks.map((t: string) => t.replace(/^\d+\.\s*/, "").trim());
      setAiSuggestions(tasks);
      setAiPrompt("");
    } catch (error) {
      alert("Failed to generate AI tasks. Please try again.");
    } finally {
      setLoadingAI(false);
    }
  };

  const handleAddAISuggestion = async (suggestion: string, index: number) => {
    try {
      await addTodo(suggestion);
      // Remove from suggestions
      setAiSuggestions(prev => prev.filter((_, i) => i !== index));
      fetchTodos();
    } catch (error) {
      alert("Failed to add task");
    }
  };

  const handleRemoveSuggestion = (index: number) => {
    setAiSuggestions(prev => prev.filter((_, i) => i !== index));
  };

  const handleLogout = () => {
    localStorage.removeItem("token");
    navigate("/login");
  };

  const completedCount = todos.filter(t => t.completed).length;
  const pendingCount = todos.filter(t => !t.completed).length;

  return (
    <div className="min-h-screen bg-gradient-to-br from-sky-100 via-blue-50 to-indigo-100 p-4 sm:p-8">
      <div className="max-w-7xl mx-auto">
        {/* Header */}
        <div className="flex items-center justify-between mb-8">
          <div className="flex items-center gap-3">
            <Zap className="w-10 h-10 text-sky-600 animate-pulse" />
            <h1 className="text-4xl sm:text-5xl font-bold text-sky-700">TaskGenius</h1>
            <Sparkles className="w-10 h-10 text-sky-600 animate-pulse" />
          </div>
          <button
            onClick={handleLogout}
            className="flex items-center gap-2 bg-sky-200/60 hover:bg-sky-300/60 text-sky-800 px-4 py-2 rounded-xl transition-all duration-300 backdrop-blur-sm"
          >
            <LogOut className="w-5 h-5" />
            <span className="hidden sm:inline">Logout</span>
          </button>
        </div>

        <p className="text-center text-sky-700 text-lg mb-8">AI-Powered Task Management</p>

        {/* Two Column Layout for Inputs */}
        <div className="grid grid-cols-1 lg:grid-cols-2 gap-6 mb-6">
          {/* Manual Task Input - LEFT */}
          <div className="bg-white/80 backdrop-blur-sm rounded-3xl shadow-xl p-6">
            <div className="flex items-center gap-2 mb-4">
              <Plus className="w-6 h-6 text-sky-600" />
              <h2 className="text-2xl font-bold text-sky-800">Add Manual Task</h2>
            </div>
            
            <div className="space-y-3">
              <input
                value={task}
                onChange={(e) => setTask(e.target.value)}
                onKeyDown={(e) => {
                  if (e.key === 'Enter') {
                    handleAdd();
                  }
                }}
                placeholder="Add a new task..."
                className="w-full p-4 border-2 border-sky-200 rounded-2xl focus:border-sky-400 focus:outline-none text-gray-700 placeholder-gray-400 transition-all duration-300"
              />
              <button
                onClick={handleAdd}
                disabled={!task.trim()}
                className="w-full bg-gradient-to-r from-sky-500 to-blue-500 text-white py-4 rounded-2xl font-semibold hover:from-sky-600 hover:to-blue-600 disabled:opacity-50 disabled:cursor-not-allowed transform hover:scale-[1.02] transition-all duration-300 shadow-lg flex items-center justify-center gap-2"
              >
                <Plus className="w-5 h-5" />
                Add Task
              </button>
            </div>
          </div>

          {/* AI Generation Card - RIGHT */}
          <div className="space-y-4">
            <div className="bg-white/80 backdrop-blur-sm rounded-3xl shadow-xl p-6">
              <div className="flex items-center gap-2 mb-4">
                <Sparkles className="w-6 h-6 text-indigo-600" />
                <h2 className="text-2xl font-bold text-indigo-800">Generate Tasks with AI</h2>
              </div>
              
              <div className="space-y-3">
                <textarea
                  value={aiPrompt}
                  onChange={(e) => setAiPrompt(e.target.value)}
                  onKeyDown={(e) => {
                    if (e.key === 'Enter' && !e.shiftKey) {
                      e.preventDefault();
                      handleGenerateAITasks();
                    }
                  }}
                  placeholder="Describe what you need to do... e.g., 'Plan a week of healthy meals' or 'Organize a team meeting'"
                  className="w-full p-4 border-2 border-indigo-200 rounded-2xl focus:border-indigo-400 focus:outline-none resize-none h-24 text-gray-700 placeholder-gray-400 transition-all duration-300"
                />
                
                <button
                  onClick={handleGenerateAITasks}
                  disabled={loadingAI || !aiPrompt.trim()}
                  className="w-full bg-gradient-to-r from-indigo-500 to-purple-500 text-white py-4 rounded-2xl font-semibold hover:from-indigo-600 hover:to-purple-600 disabled:opacity-50 disabled:cursor-not-allowed transform hover:scale-[1.02] transition-all duration-300 shadow-lg flex items-center justify-center gap-2"
                >
                  {loadingAI ? (
                    <>
                      <Loader2 className="w-5 h-5 animate-spin" />
                      Generating...
                    </>
                  ) : (
                    <>
                      <Sparkles className="w-5 h-5" />
                      Generate Tasks
                    </>
                  )}
                </button>
              </div>
            </div>

            {/* AI Suggestions Box - Appears below AI card */}
            {aiSuggestions.length > 0 && (
              <div className="bg-gradient-to-br from-purple-100 to-indigo-100 backdrop-blur-sm rounded-3xl shadow-xl p-6 border-2 border-purple-200">
                <div className="flex items-center justify-between mb-4">
                  <div className="flex items-center gap-2">
                    <Sparkles className="w-5 h-5 text-purple-600" />
                    <h3 className="text-xl font-bold text-purple-800">AI Suggestions</h3>
                  </div>
                  <span className="text-sm text-purple-600 font-medium">
                    {aiSuggestions.length} task{aiSuggestions.length !== 1 ? 's' : ''}
                  </span>
                </div>
                
                <div className="space-y-2">
                  {aiSuggestions.map((suggestion, index) => (
                    <div
                      key={index}
                      className="group bg-white/60 backdrop-blur-sm rounded-xl p-3 flex items-center gap-3 hover:bg-white/80 transition-all duration-300 border border-purple-200"
                    >
                      <span className="font-bold text-purple-600 text-sm">{index + 1}.</span>
                      <span className="flex-1 text-gray-700">{suggestion}</span>
                      <div className="flex items-center gap-2">
                        <button
                          onClick={() => handleAddAISuggestion(suggestion, index)}
                          className="p-2 text-emerald-600 hover:bg-emerald-50 rounded-lg transition-all duration-300 opacity-0 group-hover:opacity-100"
                          title="Add to Your Tasks"
                        >
                          <MoveRight className="w-5 h-5" />
                        </button>
                        <button
                          onClick={() => handleRemoveSuggestion(index)}
                          className="p-2 text-red-500 hover:bg-red-50 rounded-lg transition-all duration-300 opacity-0 group-hover:opacity-100"
                          title="Remove suggestion"
                        >
                          <X className="w-4 h-4" />
                        </button>
                      </div>
                    </div>
                  ))}
                </div>
              </div>
            )}
          </div>
        </div>

        {/* Stats */}
        <div className="grid grid-cols-3 gap-4 mb-6">
          <div className="bg-white/70 backdrop-blur-sm rounded-2xl p-4 text-center shadow-lg transform hover:scale-105 transition-all duration-300">
            <div className="text-3xl font-bold text-sky-600">{todos.length}</div>
            <div className="text-sky-700 text-sm">Total Tasks</div>
          </div>
          <div className="bg-white/70 backdrop-blur-sm rounded-2xl p-4 text-center shadow-lg transform hover:scale-105 transition-all duration-300">
            <div className="text-3xl font-bold text-emerald-600">{completedCount}</div>
            <div className="text-emerald-700 text-sm">Completed</div>
          </div>
          <div className="bg-white/70 backdrop-blur-sm rounded-2xl p-4 text-center shadow-lg transform hover:scale-105 transition-all duration-300">
            <div className="text-3xl font-bold text-blue-600">{pendingCount}</div>
            <div className="text-blue-700 text-sm">Remaining</div>
          </div>
        </div>

        {/* Todo List */}
        <div className="bg-white/80 backdrop-blur-sm rounded-3xl shadow-xl p-6">
          <div className="flex items-center gap-2 mb-6">
            <Calendar className="w-6 h-6 text-sky-600" />
            <h2 className="text-2xl font-bold text-sky-800">Your Tasks</h2>
          </div>

          <div className="space-y-3">
            {todos.length === 0 ? (
              <div className="text-center py-12 text-gray-400">
                <Plus className="w-16 h-16 mx-auto mb-3 opacity-50" />
                <p className="text-lg">No tasks yet. Add some or generate with AI!</p>
              </div>
            ) : (
              todos.map((todo) => (
                <div
                  key={todo.id}
                  className="group bg-gradient-to-r from-sky-50 to-white border-2 border-sky-100 rounded-2xl p-4 hover:border-sky-300 hover:shadow-lg transition-all duration-300 transform hover:scale-[1.01]"
                >
                  <div className="flex items-center gap-4">
                    <button
                      onClick={() => handleToggle(todo.id)}
                      className={`flex-shrink-0 w-7 h-7 rounded-full border-3 flex items-center justify-center transition-all duration-300 ${
                        todo.completed
                          ? 'bg-emerald-500 border-emerald-500'
                          : 'border-sky-300 hover:border-sky-500 hover:bg-sky-50'
                      }`}
                    >
                      {todo.completed && <Check className="w-4 h-4 text-white" />}
                    </button>

                    <div className="flex-1">
                      <p className={`text-lg font-medium transition-all duration-300 ${
                        todo.completed ? 'line-through text-gray-400' : 'text-gray-800'
                      }`}>
                        {todo.task}
                      </p>
                    </div>

                    <button
                      onClick={() => handleDelete(todo.id)}
                      className="p-2 text-gray-400 hover:text-red-500 hover:bg-red-50 rounded-xl transition-all duration-300 opacity-0 group-hover:opacity-100"
                    >
                      <Trash2 className="w-5 h-5" />
                    </button>
                  </div>
                </div>
              ))
            )}
          </div>
        </div>
      </div>
    </div>
  );
}