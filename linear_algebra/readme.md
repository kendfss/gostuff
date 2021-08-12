# [A friendly introduction to linear algebra for ML](https://youtube.com/watch?v=LlKAna21fLE)

This repository consists of lecture notes to supplement viewing of a video on the topic by Dr. Tai-Danay Bradley on the TensorFlow youtube channel.  
The subtopics covered include:  
- [Data Representation](#representation)  
- [Vector Embedding](#embedding)  
- [Dimensionality Reduction](#reduction)  

In principle, this repo expands on the lecture material by going into technical detail of implementation. I wrote this as someone with a greater than negligible acquaintance with programming who wanted to learn about ML and the mathematics that makes it possible. Also as someone familiar with Python who wants to brush up on Go and continue to live in a world where Julia is irrelevant.  


It also attempts to offer *some* generalizations of key terms and concepts. They will be given as bullet points to definitions. Feel free to get in touch if you feel I've done a particularly heinous job of that!  


There is also a [glossary](#glossary)



----------------------
    
### Representation
##### Vector
A one-dimensional array of numbers corresponding to magnitudes and indices corresponding to directions  
- A subset of the cross product of a set of Magnitudes and a set of Directions such that all terms with common direction are reducible by the binary operation which governs its dynamics  
##### Vector Space
The formal name of a set consisting of all vectors V with length N is "N-dimensional Vector Space"
##### Feature Vector
A vector whose components correspond to interesting features of an entity, as opposed to directions in a physical space
##### Feature Space
Analogue of Vector Space for feature vectors
###### Examples
- Images  
    - Consider a black and white image of width W and height M. Each of its pixels can be represented by a 0 or 1 in a 2d-matrix which, in turn, can be unwrapped by tracing each of its rows from, for instance, left to right  
- Words in Documents
    - Given a collection of N documents, you can create a vector for each word from every document counting the number of times it appears in the nth document  
    - The vectors can be assembled into a matrix which can be useful for [*Latent Semantic Analysis*]()
- One Hot Encodings
    - Given an ordered collection of N words, each one can be assigned a unique vector indicating its position in that collection. These vectors would consist of N items; a 1 indicating the position of the vector in the collection, a 0 elsewhere.  
        - These are formally known as [Standard Basis Vectors]()  
        - These can be very [sparse](#sparse)  
        - Possible lack o meaningful relationships between vectors  
            - when similarity is evaluated via [dot products](), One-hots imply complete dis-similarity  
### Embedding
An embedding of a vector V0 is vector with lower dimension V1 that can represent the information of V0 without losing fidelity  
###### Matrix Factorization
A matrix is a 2d array of numbers that can be associated with the process of transforming one vector into another (stretching, scaling, or something more complex). That all matrices can be factored is a fundamental theorem of linear algebra. The process of matrix factorization yields a [Single Value Decomposition]() of the matrix.  
![]( Matrix IO image)
**Singular Value Decomposition**  
    - Any Matrix can be written as the product of three smaller matrices  
    ![]( SVD image)
    - see [The Extraordinary SVD]() (Martin & Porter 2012)

    
###### Neural Networks
### Reduction
**Principal component analysis** Is the art of uncovering a matrix' eigenvalues. An eigenvalue E of a matrix M is a vector whose cross product with M, V, is a scalar multiple of E. The scalar of V is called an eigenvalue of V.
### Glossary
###### Sparse
- A Vector is sparse if it contains lots of zeros  
    - A vector is sparse if a shorter vector can be used to encode its essential information  
