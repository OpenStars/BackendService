
namespace cpp OpenStars.Core.IDGenerate
namespace java openstars.core.idgenerate
namespace go openstars.core.idgen

exception InvalidOperation {
  1: i32 error,
  2: string message
}

service TGenerator {
	i32 createGenerator(1:string genName)
		throws (1:InvalidOperation ouch),
		
	i32 removeGenerator(1:string genName)
		throws (1:InvalidOperation ouch),
		
	i64 getCurrentValue(1:string genName)
		throws (1:InvalidOperation ouch),
		
	i64 getValue(1:string genName)
		throws (1:InvalidOperation ouch),

	i64 getStepValue(1:string genName,2:i64 step)
		throws (1:InvalidOperation ouch),
}
