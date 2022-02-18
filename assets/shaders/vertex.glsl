#version 440

layout (location = 0) in vec3 vertex_position;
layout (location = 1) in vec3 vertex_color;

uniform mat4 projection;
uniform mat4 camera;
uniform mat4 model;

out vec3 vs_position;
out vec3 vs_color;

void main() {
    vs_position = vec4(model * vec4(vertex_position, 1.0)).xyz;
    vs_color = vertex_color;

    gl_Position = projection * camera * model * vec4(vertex_position, 1.0);
}
