import { useState } from "react";


interface Task {
    title: string;
    description: string;
    completed?: boolean;
    category?: string;
    dueDate?: string;
    priority?: string
}
const ToDoList = () => {

    const categories = ["Work", "Personal", "Shopping", "Others"];
    const priorities = ["High", "Medium", "Low"];

    const handleSubmit = (e: React.FormEvent) => {
        e.preventDefault();
        setTasks((prev) => {
            const updatedTasks = [...prev, { title, description, dueDate, category, priority }];
            console.log(updatedTasks);
            return updatedTasks;
        });

        setTitle('');
        setDescription('');
        setDueDate("");
    }

    const handleDelete = (id: number) => {
        console.log("Delete", id);

        setTasks((prev) => {
            const updateTasks = prev.filter((_, index) => index !== id)
            return updateTasks;
        })
    }
    const [tasks, setTasks] = useState<Task[]>([]);
    const [title, setTitle] = useState<string>('');
    const [description, setDescription] = useState<string>('');
    const [category, setCategory] = useState<string>('');
    const [dueDate, setDueDate] = useState<string>('');
    const [priority, setPriority] = useState<string>('');
    const [completed, setCompleted] = useState<boolean>(false);
    return (
        <div className="w-5/12 mx-auto  mt-10 p-6 bg-white rounded-lg shadow-lg">
            <h1 className="text-2xl font-bold text-center mb-4">📝 To-Do List</h1>
            <form onSubmit={handleSubmit} className="flex flex-col gap-4">
                <input
                    type="text"
                    placeholder="Enter your title"
                    className="p-2 border rounded focus:outline-blue-500"
                    value={title}
                    onChange={(e) => setTitle(e.target.value)}
                />
                <textarea
                    rows={4}
                    placeholder="Enter your description"
                    className="p-2 border rounded focus:outline-blue-500"
                    value={description}
                    onChange={(e) => setDescription(e.target.value)}
                ></textarea>

                {/* 📂 Category Dropdown */}
                <select
                    className="p-2 border rounded focus:outline-blue-500"
                    value={category}
                    onChange={(e) => setCategory(e.target.value)}
                >
                    <option value="">Select Category</option>
                    {categories.map((val, index) => (
                        <option key={index} value={val}>
                            {val}
                        </option>
                    ))}
                </select>

                {/* 🔥 Priority Dropdown */}
                <select
                    className="p-2 border rounded focus:outline-blue-500"
                    value={priority}
                    onChange={(e) => setPriority(e.target.value)}
                >
                    <option value="">Select Priority</option>
                    {priorities.map((val, index) => (
                        <option key={index} value={val}>
                            {val}
                        </option>
                    ))}
                </select>

                {/* 📅 Date Input */}
                <input
                    type="date"
                    className="p-2 border rounded focus:outline-blue-500"
                    value={dueDate}
                    onChange={(e) => setDueDate(e.target.value)}
                />

                <button
                    type="submit"
                    className="bg-blue-500 text-white py-2 rounded hover:bg-blue-600 transition duration-200"
                >
                    Add Task
                </button>
            </form>

            <ul className="mt-6 space-y-4">
                {tasks.map((task, index) => (
                    <li
                        key={index}
                        className={`relative p-5 rounded-lg shadow-md transition duration-200 flex justify-between items-start border-l-8 ${task.priority === "High"
                            ? "border-red-500 bg-red-50"
                            : task.priority === "Medium"
                                ? "border-yellow-500 bg-yellow-50"
                                : "border-green-500 bg-green-50"
                            }`}
                    >
                        {/* Task Content */}
                        <div className="flex flex-col w-full">
                            <h3
                                className={`text-lg font-bold ${task.completed ? "line-through text-gray-400" : "text-gray-800"
                                    }`}
                            >
                                {task.title}
                            </h3>
                            <p className="text-gray-600 text-sm">{task.description}</p>

                            {/* Task Meta Info */}
                            <div className="flex flex-wrap gap-2 mt-2 text-sm">
                                <span className="flex items-center bg-gray-200 text-gray-700 px-2 py-1 rounded">
                                    📅 {task.dueDate}
                                </span>
                                <span
                                    className={`px-2 py-1 rounded text-white ${task.priority === "High"
                                        ? "bg-red-500"
                                        : task.priority === "Medium"
                                            ? "bg-yellow-500 text-black"
                                            : "bg-green-500"
                                        }`}
                                >
                                    {task.priority} Priority
                                </span>
                                <span className="bg-blue-100 text-blue-600 px-2 py-1 rounded">
                                    📂 {task.category}
                                </span>
                            </div>

                            {/* Mark as Completed Checkbox */}
                            <label className="flex items-center mt-2 cursor-pointer">
                                <input
                                    type="checkbox"
                                    checked={task.completed}
                                    onChange={() => setCompleted(!completed)}
                                    className="mr-2"
                                />
                                <span className="text-gray-600 text-sm">
                                    {task.completed ? "Completed" : "Mark as Done"}
                                </span>
                            </label>
                        </div>

                        {/* Delete Button */}
                        <button
                            onClick={() => handleDelete(index)}
                            className="absolute top-3 right-3 text-white p-1.5 rounded-full hover:transition duration-200"
                        >
                            ✖
                        </button>
                    </li>
                ))}
            </ul>

        </div>
    )
}

export default ToDoList;