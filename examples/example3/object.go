package main

import (
	"github.com/o5h/glm"
	"github.com/o5h/glx/material"
	"github.com/o5h/glx/mesh"

	"github.com/o5h/glx/renderer"
)

type Object struct {
	Transform glm.Transform
	Mesh      *mesh.Mesh
	Material  *material.Instance

	m glm.Mat4x4
}

func NewObject(mesh *mesh.Mesh, material *material.Instance) *Object {
	o := &Object{}
	o.Transform.Scale.SetXYZ(1, 1, 1)
	o.Mesh = mesh
	o.Material = material
	return o
}

func (o *Object) Draw() {
	o.m.SetTransform(&o.Transform.Position,
		&o.Transform.Scale,
		&o.Transform.Rotation)
	renderer.SetModel(&o.m)
	renderer.UseMaterial(o.Material)
	renderer.DrawMesh(o.Mesh)
}
