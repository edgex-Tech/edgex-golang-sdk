from setuptools import setup, find_packages

setup(
    name="edgex-python-sdk",
    version="0.1.0",
    description="A Python SDK for interacting with the EdgeX Exchange API",
    author="EdgeX Tech",
    author_email="info@edgex.exchange",
    packages=find_packages(),
    install_requires=[
        "requests>=2.25.0",
        "websocket-client>=1.0.0",
        "pydantic>=1.8.0",
        "python-dotenv>=0.15.0",
        "pycryptodome>=3.15.0",
        "ecdsa>=0.17.0",
    ],
    classifiers=[
        "Development Status :: 3 - Alpha",
        "Intended Audience :: Developers",
        "License :: OSI Approved :: MIT License",
        "Programming Language :: Python :: 3",
        "Programming Language :: Python :: 3.7",
        "Programming Language :: Python :: 3.8",
        "Programming Language :: Python :: 3.9",
    ],
    python_requires=">=3.7",
)