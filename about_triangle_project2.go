package about_triangle_project2

import "./koans"
import "./triangle"

func TestIllegalTrianglesReturnErrors(t *koans.T) {
	_, err := triangle.Triangle(0, 0, 0)
	t.AssertTrue(err != nil)
	_, err = triangle.Triangle(3, 4, -5)
	t.AssertTrue(err != nil)
	_, err = triangle.Triangle(1, 1, 3)
	t.AssertTrue(err != nil)
	_, err = triangle.Triangle(2, 4, 2)
	t.AssertTrue(err != nil)
}
