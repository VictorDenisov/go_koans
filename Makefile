RUNNERFILES=\
	koan_runner.go
KOANSFILES=\
	about_asserts.go\
	about_nil.go\
	about_arrays.go\
	about_strings.go\
	about_tuples.go\
	about_methods.go\
	about_control_statements.go\
	about_triangle_project.go\
	about_triangle_project2.go\
	triangle.go\
	about_closures.go\
	about_maps.go
KOANSLIB=\
	koans.go
SOURCES=$(KOANSLIB) $(KOANSFILES) $(RUNNERFILES)
OBJECTS=$(SOURCES:.go=.8)


%.8:%.go
	8g $<

all: koan_runner
	./koan_runner

about_triangle_project.8: triangle.8

koan_runner: $(OBJECTS)
	8l -o koan_runner koan_runner.8

clean:
	rm -f $(OBJECTS)
	

