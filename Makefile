include $(GOROOT)/src/Make.inc

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
	about_scoring_project.go\
	about_triangle_project.go\
	about_triangle_project2.go\
	about_structs.go\
	about_dice_project.go\
	triangle.go\
	about_closures.go\
	about_channels.go\
	about_interfaces.go\
	about_embedding.go\
	about_maps.go
KOANSLIB=\
	koans.go
SOURCES=$(KOANSLIB) $(KOANSFILES) $(RUNNERFILES)
OBJECTS=$(SOURCES:.go=.${O})


%.${O}:%.go
	${O}g $<

all: koan_runner
	./koan_runner

about_triangle_project.${O}: triangle.${O}

koan_runner: $(OBJECTS)
	${O}l -o koan_runner koan_runner.${O}

clean:
	rm -f $(OBJECTS)

